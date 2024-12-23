package handlers

import (
	"boobook-chat-microservice/internal/http/requests"
	"boobook-chat-microservice/internal/repositories"
	"boobook-chat-microservice/internal/services"
	"boobook-chat-microservice/internal/slogger"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type messageHandler struct {
	logger         *slog.Logger
	messageService services.MessageService
}

func NewMessageHandler(logger *slog.Logger, messageService services.MessageService) MessageHandler {
	return &messageHandler{
		logger:         logger,
		messageService: messageService,
	}
}

func (h *messageHandler) Create(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var SendMessageRequest requests.SendMessageRequest
	if err = ctx.BindJSON(&SendMessageRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload: " + err.Error()})
		return
	}

	SendMessageRequest.ChatId = id
	userId, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	SendMessageRequest.UserId = (userId).(int)

	if err = h.messageService.Create(&SendMessageRequest); err != nil {
		// Если пользователь не является участником чата, возвращаем ошибку
		if errors.Is(err, repositories.ErrUserNotMember) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "user is not a member of the chat"})
			return
		}

		// В случае других ошибок возвращаем 400
		h.logger.ErrorContext(ctx, "failed to send message", slogger.ErrorToSlogAttr(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}

func (h *messageHandler) GetList(ctx *gin.Context) {
	// Получаем id чата из параметров запроса
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// Получаем id пользователя из контекста
	userId, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	messages, err := h.messageService.GetList(id, (userId).(int))
	if err != nil {
		// Если пользователь не является участником чата, возвращаем ошибку
		if errors.Is(err, repositories.ErrUserNotMember) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "user is not a member of the chat"})
			return
		}

		// В случае других ошибок возвращаем 400
		h.logger.ErrorContext(ctx, "failed to get messages", slogger.ErrorToSlogAttr(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messages": messages})
}
