package handlers

import (
	"intern/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRole godoc
// @Router       /api/role [POST]
// @Security BearerAuth
// @Summary      Create Role
// @Description  Create a new role
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        role body domain.RoleCreateRequest true "Role Request"
// @Success      200  {object}  domain.ResponseSuccess
// @Failure      400  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) CreateRole(c *gin.Context) {
	var (
		payload domain.RoleCreateRequest
	)
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()

	err := c.ShouldBindJSON(&payload)
	if h.handleError(c, err) {
		return
	}

	err = h.services.CreateRole(ctxWithTimeout, &payload)
	if h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, domain.ResponseSuccess{
		Data:   "role created successfully",
		Status: http.StatusOK,
	})
}

// GetRoles godoc
// @Router       /api/roles [GET]
// @Security BearerAuth
// @Summary      List Roles
// @Description  Get all active roles
// @Tags         Role
// @Produce      json
// @Success      200  {object}  domain.RoleListResponse
// @Failure      400  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) GetRoles(c *gin.Context) {
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()

	roles, err := h.services.GetRoleList(ctxWithTimeout)
	if h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, roles)
}

// UpdateRole godoc
// @Router       /api/role/{id} [PUT]
// @Security BearerAuth
// @Summary      Update Role
// @Description  Update an existing role's name
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        id path string true "Role ID"
// @Param        role body domain.RoleCreateRequest true "Role Request"
// @Success      200  {object}  domain.ResponseSuccess
// @Failure      400  {object}  domain.ResponseError
// @Failure      404  {object}  domain.ResponseError
// @Failure      500  {object}  domain.ResponseError
func (h *Handler) UpdateRole(c *gin.Context) {
	var payload domain.RoleCreateRequest
	ctxWithTimeout, cancel := h.makeContext()
	defer cancel()

	roleID := c.Param("id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, domain.ResponseError{Message: "role id is required", Status: http.StatusBadRequest})
		return
	}

	if err := c.ShouldBindJSON(&payload); h.handleError(c, err) {
		return
	}

	if err := h.services.UpdateRole(ctxWithTimeout, roleID, payload.Name); h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, domain.ResponseSuccess{Data: "role updated successfully", Status: http.StatusOK})
}
