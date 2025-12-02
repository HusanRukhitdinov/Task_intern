package handlers

import (
	"context"
	"fmt"
	configs2 "intern/internal/configs"
	"intern/internal/core/repository"
	"intern/internal/core/services"
	"intern/pkg/email"
	"intern/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Run()
	Stop()
}

type Handler struct {
	engine   *gin.Engine
	services *services.Service
	log      logger.ILogger
	cfg      configs2.Config
}

func NewHandler(engine *gin.Engine, services *services.Service, log logger.ILogger, cfg configs2.Config) *Handler {
	return &Handler{
		engine:   engine,
		services: services,
		log:      log,
		cfg:      cfg,
	}
}

func NewServer(cfg configs2.Config) Server {

	loggerLevel := logger.LevelDebug
	switch cfg.Environment {
	case configs2.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case configs2.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}
	log := logger.NewLogger("intern.uz", loggerLevel)
	defer logger.Cleanup(log)

	engine := gin.Default()
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowCredentials = true
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders,
		"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,"+
			"Authorization, accept, origin, Cache-Control, X-Requested-With")
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders, "*")
	defaultConfig.AllowMethods = append(defaultConfig.AllowMethods, "OPTIONS")

	engine.Use(cors.New(defaultConfig))

	dbTx, _ := repository.New(context.Background(), cfg, log)

	// Initialize email sender
	emailSender := email.NewEmailSender(email.EmailConfig{
		SMTPHost:     cfg.SMTPHost,
		SMTPPort:     cfg.SMTPPort,
		SMTPUsername: cfg.SMTPUsername,
		SMTPPassword: cfg.SMTPPassword,
		FromEmail:    cfg.FromEmail,
		FromName:     cfg.FromName,
	})

	services1 := services.NewService(dbTx, log, emailSender)

	handler := NewHandler(engine, services1, log, cfg)
	setUpApi(handler)

	return handler
}

// @title intern.UZ App
// @description This API contains the source for the intern.uz app
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /v1

// Run initializes http server
// Run initializes http server
func (h *Handler) Run() {
	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL(fmt.Sprintf(
			"%s/%d/swagger/docs.json",
			h.cfg.HTTHost,
			h.cfg.HTTPPort,
		)),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

	h.log.Info("server is running: ", logger.Any("address", fmt.Sprintf("%s:%d", h.cfg.HTTHost, h.cfg.HTTPPort)))
	h.log.Info("swagger: ", logger.Any("url", fmt.Sprintf("http://%s:%d/swagger/index.html", h.cfg.HTTHost, h.cfg.HTTPPort)))

	if err := h.engine.Run(fmt.Sprintf(":%d", h.cfg.HTTPPort)); err != nil {
		h.log.Error("failed to run server", logger.Error(err))
	}

}

func (h *Handler) Stop() {
	h.log.Info("shutting down")
}
