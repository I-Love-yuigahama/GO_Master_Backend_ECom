package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Ecom.com/TUT/internal/config"
	"Ecom.com/TUT/internal/database"
	"Ecom.com/TUT/internal/events"
	"Ecom.com/TUT/internal/interfaces"
	"Ecom.com/TUT/internal/logger"
	"Ecom.com/TUT/internal/providers"
	"Ecom.com/TUT/internal/services"
	"Ecom.com/TUT/server"
	"github.com/gin-gonic/gin"
)

// @title E-Commerce API
// @version 1.0
// @description A modern e-commerce API built with Go, Gin, and GORM
// @termsOfService http://swagger.io/terms/

// @contact.name   Joseph Abah
// @contact.url    http://linkedin.com/in/abahjoseph
// @contact.email  no-email@no-email

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @schemas http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	log := logger.New()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load db")
	}

	mainDB, err := db.DB()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to get DB connection")
	}

	defer mainDB.Close()

	ctx := context.Background()

	eventPublisher , err := events.NewEventPublisher(ctx, &cfg.AWS)
	if err != nil {
	log.Error().Err(err).Msg("failed to create event publisher")
		return
	}

	gin.SetMode(cfg.Server.GinMode)



	authServices := services.NewAuthService(db, cfg,eventPublisher)
	productService := services.NewProductServices(db)
	userService := services.NewUserService(db)
	cartService := services.NewCartServices(db)
	ordertService := services.NewOderderSerice(db)

	var uploadProvider interfaces.UploadProvider

	if cfg.Upload.UploadProvider == "s3" {
		uploadProvider = providers.NewS3Provider(cfg)
	} else {
		uploadProvider = providers.NewLocalUploadProvider(cfg.Upload.Path)
	}

	uploadService := services.NewUploadServices(uploadProvider)

	srv := server.New(cfg, db, &log, authServices, productService, userService, uploadService, cartService, ordertService)
	router := srv.SetupRoutes()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Str("port", cfg.Server.Port).Msg("starting http server")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("failed to start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server")
		return
	}

	log.Info().Msg("shuting down http server")

}
