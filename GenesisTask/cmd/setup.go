package main

import (
	"GenesisTask/config"
	"GenesisTask/pkg/application"
	"GenesisTask/pkg/delivery/handlers"
	"GenesisTask/pkg/infrastructure/crypto"
	"GenesisTask/pkg/infrastructure/email"
	cache "GenesisTask/pkg/infrastructure/storage/cache/redis"
	storage "GenesisTask/pkg/infrastructure/storage/subscription_repository"
	"time"

	"github.com/gin-gonic/gin"
)

func LaunchEngine(logger application.Logger) {
	gin.SetMode(gin.ReleaseMode)
	logger.LogDebug("Set engine mode to release")
	router := gin.Default()
	handlers := createHandlers(logger)
	logger.LogDebug("Created handlers")
	initRoutes(router, handlers)
	logger.LogInfo("Routes configured")
	logger.LogDebug("Running on " + config.ServerUrl)
	router.Run(config.ServerUrl)
}

func initRoutes(router *gin.Engine, h *handlers.Handlers) {
	router.POST("/api/rate", h.GetRate)
	router.POST("/api/subscribe", h.Subscribe)
	router.POST("/api/sendEmails", h.SendEmailsToUsers)
}

func createHandlers(logger application.Logger) (h *handlers.Handlers) {
	logger.LogDebug("Creating new providers chain")
	providersChain := crypto.NewProvidersChain(logger)
	emailSender := email.NewGomailSender(logger)
	logger.LogDebug("Created new gomail sender")
	cacheDuration := time.Duration(config.CacheDurationMins) * time.Minute
	cache := cache.NewRedisCache(config.CacheDb, cacheDuration)
	logger.LogDebug("Created new redis cache")
	emailAddressesStorage := storage.NewSubscriptionFileRepository(logger)
	logger.LogDebug("Created new subscription file repository")

	r1 := application.NewRateRepository(*providersChain, cache, logger)
	r2 := application.NewSubscriptionRepository(emailAddressesStorage, logger)
	r3 := application.NewEmailSenderRepository(emailAddressesStorage, emailSender, *r1, logger)
	logger.LogDebug("Initialized repositories")

	h1 := handlers.NewRateHandler(r1, logger)
	h2 := handlers.NewSubscribeHandler(*r2, logger)
	h3 := handlers.NewSendRateEmailsHandler(*r3, logger)
	logger.LogDebug("Initialized handlers")

	return handlers.NewHandlers(h1, h2, h3)
}
