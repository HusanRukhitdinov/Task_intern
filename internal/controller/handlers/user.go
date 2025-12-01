package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"intern/internal/core/domain"
	token "intern/pkg/jwt"
	"intern/pkg/middleware"
	"intern/pkg/validate"
)

// RegisterUser   godoc
// @Router       /api/register/user [POST]
// @Security     ApiKeyAuth
// @Summary      User Create
// @Description  User Create
// @Tags         User
// @Produce      json
// @Param        user body domain.RegisterRequest true "User Request"
// @Success      200  {object}  domain.TokenResponse
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) RegisterUser(c *gin.Context) {
	var (
		payload domain.RegisterRequest
	)
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()
	err := c.ShouldBindJSON(&payload)
	if h.handleError(c, err) {
		return
	}
	err = validate.CheckEmailAndPassword(payload.Email)

	if err != nil {
		c.JSON(400, domain.ResponseSuccess{
			Data:   "email is not strong",
			Status: 400,
		})
		return
	}
	if !validate.IsValidPassword(payload.HashPassword) {
		c.JSON(400, domain.ResponseSuccess{
			Data:   "Password is not valid",
			Status: 400,
		})
		return
	}
	hashedPassword, err := h.hashingPassword([]byte(payload.HashPassword))
	if h.handleError(c, err) {
		return
	}
	payload.HashPassword = hashedPassword

	err = h.services.RegisterUser(ctxWithTimeout, &payload)

	if h.handleError(c, err) {
		return
	}

	c.JSON(200, domain.ResponseSuccess{
		Data:   "success register on user",
		Status: 200,
	})

}

// UpdateUser   godoc
// @Router       /api/user [put]
// @Security BearerAuth
// @Summary      User Update
// @Description  User Update
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body domain.UserParams true "User Request"
// @Success      200  {object}  domain.ResponseError
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) UpdateUser(c *gin.Context) {
	var (
		payload domain.UserParams
	)
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()
	err := c.ShouldBindJSON(&payload)
	if h.handleError(c, err) {
		return
	}

	userID, err := middleware.GetIdByClaims(c, h.log)
	if h.handleError(c, err) {
		return
	}
	payload.ID = userID
	err = h.services.EditOneUser(ctxWithTimeout, &payload)
	if h.handleError(c, err) {
		return
	}
	c.JSON(200, domain.ResponseSuccess{
		Data:   "success update on user",
		Status: 200,
	})

}

// DeleteUser   godoc
// @Router       /api/user/{id} [delete]
// @Security BearerAuth
// @Summary      User Delete
// @Description  User Delete
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id path domain.IDRequest true "User ID"
// @Success      200  {object}  string
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) DeleteUser(c *gin.Context) {
	var (
		id string
	)
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()
	id = c.Param("id")
	err := h.services.DeleteOneUser(ctxWithTimeout, &domain.IDRequest{
		ID: id,
	})
	if h.handleError(c, err) {
		return
	}

	c.JSON(200, domain.ResponseSuccess{
		Data:   "success deleted on user",
		Status: 200,
	})

}

// GetUserAllUsers godoc
// @Security     ApiKeyAuth
// @Router       /api/users [GET]
// @Summary      Get all admins
// @Description  get all admins
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        filter query domain.GetListRequest false "filter"
// @Success      200  {object}  domain.SelectManyUsers
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) GetUserAllUsers(c *gin.Context) {
	response, err := h.services.SelectManyUsers(context.Background())
	if h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, response)

}

// GetUser godoc
// @Security     ApiKeyAuth
// @Router       /api/user [GET]
// @Summary      Get user  by id
// @Description  get user  by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.SelectOneUser
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) GetUser(c *gin.Context) {
	var (
		payload domain.IDRequest
	)
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()
	userID, err := middleware.GetIdByClaims(c, h.log)
	if h.handleError(c, err) {
		return
	}

	payload.ID = userID
	admin, err := h.services.SelectOneUser(ctxWithTimeout, &payload)
	if h.handleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, admin)

}

// LoginUser godoc
// @Security     ApiKeyAuth
// @Router       /api/login/user [POST]
// @Summary      Get user  by id
// @Description  get user   by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body domain.LoginRequest true "User Request"
// @Success      200  {object}  domain.TokenResponse
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) LoginUser(c *gin.Context) {
	var payload domain.LoginRequest
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()
	err := c.ShouldBindJSON(&payload)
	if h.handleError(c, err) {
		return
	}

	resp, err := h.services.SelectOneUserEmail(ctxWithTimeout, &domain.EmailRequest{
		Email: payload.Email,
	})
	if h.handleError(c, err) {
		return
	}

	err = h.compareHashedPassword([]byte(resp.Password), []byte(payload.HashPassword))
	if h.handleError(c, err) {
		return
	}

	jwtToken, err := token.GenerateJWTToken(&domain.TokenRequest{
		ID:   resp.ID,
		Role: "admin",
	}, h.log)
	if h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jwtToken)

}
