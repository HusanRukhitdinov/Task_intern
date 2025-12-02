package handlers

import (
	_ "intern/api/openapi"
	"intern/internal/controller/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setUpApi(h *Handler) {

	v1 := h.engine.Group("/api")

	v1.POST("/otp", h.SendOTP)
	v1.POST("/otp/confirm", h.ConfirmOTP)
	v1.POST("/signup", h.Signup)
	v1.POST("/login", h.Login)


	v1.Use(middleware.AuthorizerMiddleware(h.cfg))
	{
		v1.POST("/role", h.CreateRole)
		v1.GET("/roles", h.GetRoles)
		v1.POST("/sysuser", h.CreateSysuser)
	}

	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
