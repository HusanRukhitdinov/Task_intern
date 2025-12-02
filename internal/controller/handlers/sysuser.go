package handlers

import (
	"intern/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateSysuser godoc
// @Router       /api/sysuser [POST]
// @Security BearerAuth
// @Summary      Create Sysuser
// @Description  Create a new system user with roles
// @Tags         Sysuser
// @Accept       json
// @Produce      json
// @Param        sysuser body domain.SysuserCreateRequest true "Sysuser Request"
// @Success      200  {object}  domain.SysuserCreateResponse
// @Failure      400  {object}  domain.ResponseError
// @Failure      409  {object}  domain.ResponseError "Sysuser already exists or role not found"
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) CreateSysuser(c *gin.Context) {
	var payload domain.SysuserCreateRequest

	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()

	err := c.ShouldBindJSON(&payload)
	if h.handleError(c, err) {
		return
	}
	hashedPassword, err := h.hashingPassword([]byte(payload.Password))
	if h.handleError(c, err) {
		return
	}

	createdBy := ""

	response, err := h.services.CreateSysuser(ctxWithTimeout, &payload, hashedPassword, createdBy)

	if err != nil {
		if err == domain.ErrSysuserAlreadyExists {
			c.JSON(http.StatusConflict, domain.ResponseError{
				Message: "sysuser already exists",
				Status:  http.StatusConflict,
			})
			return
		}
		if err == domain.ErrRoleNotFound {
			c.JSON(http.StatusConflict, domain.ResponseError{
				Message: "role not found",
				Status:  http.StatusConflict,
			})
			return
		}
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
