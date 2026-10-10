package main

import (
	"Ecom.com/TUT/internal/config"
	"Ecom.com/TUT/internal/database"
	"Ecom.com/TUT/internal/logger"
	"github.com/gin-gonic/gin"
)

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

	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("staring server")
}
