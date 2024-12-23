package main

import (
	"boobook-chat-microservice/internal/config"
	"boobook-chat-microservice/internal/http/handlers"
	"boobook-chat-microservice/internal/http/router"
	"boobook-chat-microservice/internal/repositories/postgres"
	"boobook-chat-microservice/internal/services"
	"boobook-chat-microservice/internal/slogger"
	"database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	logger := slogger.NewLogger()
	cfg := config.MustLoad()

	dbConnection, err := postgres.NewConnection(cfg.DBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	defer func(dbConn *sql.DB) {
		if err = postgres.CloseConnection(dbConn); err != nil {
			logger.Error("failed to close the database connection", slogger.ErrorToSlogAttr(err))
			return
		}
	}(dbConnection)

	messageRepository := postgres.NewMessageRepository(dbConnection)
	messageService := services.NewMessageService(messageRepository)
	messageHandler := handlers.NewMessageHandler(logger, messageService)

	// Setup server
	httpServer := &http.Server{
		Addr:           ":" + cfg.HTTPServerConfig.Port,
		Handler:        router.SetupRouter(logger, messageHandler),
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    cfg.HTTPServerConfig.Timeout * time.Second,
		WriteTimeout:   cfg.HTTPServerConfig.Timeout * time.Second,
	}

	// Start server
	if err = httpServer.ListenAndServe(); err != nil {
		panic(fmt.Errorf("failed to start the server: %w", err))
	}
}
