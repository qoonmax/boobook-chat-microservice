package router

import (
	"boobook-chat-microservice/internal/http/handlers"
	"boobook-chat-microservice/internal/http/middlewares"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func SetupRouter(
	logger *slog.Logger,
	messageHandler handlers.MessageHandler,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.Use(middlewares.Log(logger))

		users := api.Group("/chats")
		{
			users.Use(middlewares.Auth()).POST("/:id/send-message", messageHandler.Create)
			users.Use(middlewares.Auth()).GET("/:id/list-messages", messageHandler.GetList)
		}
	}

	return router
}
