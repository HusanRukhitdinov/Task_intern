package handlers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "intern/api/openapi"
	"intern/pkg/middleware"
)

func setUpApi(h *Handler) {

	v1 := h.engine.Group("/api")

	v1.POST("/login/user", h.LoginUser)
	v1.POST("/register/user", h.RegisterUser)

	v1.Use(middleware.JWTMiddleware())

	//v1.Use(casbin.CasbinMiddleware(h.enforcer))

	v1.Use()
	{
		v1.POST("/answer", h.CreateAnswer)
		v1.PUT("/answer/:id", h.UpdateAnswer)
		v1.DELETE("/answer/:id", h.DeleteAnswer)
		v1.GET("/answer/:id", h.GetOneAnswer)
		v1.GET("/answers/:id", h.GetAllAnswersQuestionID)

		v1.POST("/check/question", h.CheckQuestion)

		v1.POST("/question", h.CreateQuestion)
		v1.PUT("/question/:id", h.UpdateQuestion)
		v1.DELETE("/question/:id", h.DeleteQuestion)
		v1.GET("/questions/:id", h.GetAllQuestionsinternID)

		v1.POST("/intern", h.Createintern)
		v1.PUT("/intern/:id", h.Updateintern)
		v1.PUT("/intern/:id/difficulty", h.UpdateinternDifficulty)
		v1.DELETE("/intern/:id", h.Deleteintern)
		v1.GET("/intern/:id", h.Getintern)
		v1.GET("/internzes", h.GetAllinternzes)

		v1.POST("/intern/:id/start", h.Startintern)
		v1.GET("/intern/:id/result", h.GetinternResult)
		v1.PUT("/intern/:id/completed/:intern_id", h.Completedintern)

		v1.GET("/user", h.GetUser)
		v1.DELETE("/user/:id", h.DeleteUser)
		v1.GET("/users", h.GetUserAllUsers)

		v1.PUT("/user", h.UpdateUser)

	}

	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
