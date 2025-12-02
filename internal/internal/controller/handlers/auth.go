package handlers

import (
	"net/http"

	"intern/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// SendOTP godoc
// @Router /api/otp [POST]
// @Security BearerAuth
// @Summary Send OTP
// @Description Send a one-time password to the given email
// @Tags OTP
// @Accept json
// @Produce json
// @Param otp body domain.OTPCreateRequest true "OTP request"
// @Success 200 {object} domain.OTPResponse
// @Failure 400 {object} domain.ResponseError
// @Failure 500 {object} domain.ResponseError
func (h *Handler) SendOTP(c *gin.Context) {
	var req domain.OTPCreateRequest
	if err := c.ShouldBindJSON(&req); h.handleError(c, err) {
		return
	}

	resp, err := h.services.SendOTP(c.Request.Context(), &req)
	if h.handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ConfirmOTP godoc
// @Router /api/otp/confirm [POST]
// @Security BearerAuth
// @Summary Confirm OTP
// @Description Validate the OTP code and receive a confirmation token
// @Tags OTP
// @Accept json
// @Produce json
// @Param otp body domain.OTPConfirmRequest true "OTP confirm request"
// @Success 200 {object} domain.OTPConfirmationTokenResponse
// @Failure 400 {object} domain.ResponseError
// @Failure 401 {object} domain.ResponseError
// @Failure 500 {object} domain.ResponseError
func (h *Handler) ConfirmOTP(c *gin.Context) {
	var req domain.OTPConfirmRequest
	if err := c.ShouldBindJSON(&req); h.handleError(c, err) {
		return
	}

	resp, err := h.services.ConfirmOTP(c.Request.Context(), &req)
	if h.handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Signup godoc
// @Router /api/signup [POST]
// @Security BearerAuth
// @Summary Signup (after OTP confirmation)
// @Tags Auth
// @Accept json
// @Produce json
// @Param signup body domain.SignupRequest true "Signup request"
// @Success 200 {object} domain.SignupResponse
// @Failure 400 {object} domain.ResponseError
// @Failure 401 {object} domain.ResponseError
// @Failure 409 {object} domain.ResponseError
func (h *Handler) Signup(c *gin.Context) {
	var req domain.SignupRequest
	if err := c.ShouldBindJSON(&req); h.handleError(c, err) {
		return
	}

	resp, err := h.services.Signup(c.Request.Context(), &req)
	if err != nil {
		switch err {
		case domain.ErrOTPConfirmationTokenExpired:
			c.JSON(http.StatusUnauthorized, domain.ResponseError{Message: err.Error(), Status: http.StatusUnauthorized})
		case domain.ErrUnavailableOtpConfirmation:
			c.JSON(http.StatusBadRequest, domain.ResponseError{Message: err.Error(), Status: http.StatusBadRequest})
		case domain.ErrUserAlreadyExists:
			c.JSON(http.StatusConflict, domain.ResponseError{Message: err.Error(), Status: http.StatusConflict})
		default:
			h.handleError(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Login godoc
// @Router /api/login [POST]
// @Security ApiKeyAuth
// @Summary Login (user or sysuser)
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "Login request"
// @Success 200 {object} domain.TokenResponse
// @Failure 400 {object} domain.ResponseError
// @Failure 401 {object} domain.ResponseError
func (h *Handler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); h.handleError(c, err) {
		return
	}

	resp, err := h.services.Login(c.Request.Context(), &req)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, domain.ResponseError{Message: err.Error(), Status: http.StatusUnauthorized})
			return
		}
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
