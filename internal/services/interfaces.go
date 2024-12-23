package services

import (
	"boobook-chat-microservice/internal/http/requests"
	"boobook-chat-microservice/internal/models"
)

type MessageService interface {
	Create(sendMessageRequest *requests.SendMessageRequest) error
	GetList(chatId int, userId int) ([]*models.Message, error)
}
