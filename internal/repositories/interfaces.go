package repositories

import (
	"boobook-chat-microservice/internal/models"
)

type MessageRepository interface {
	Create(chatId int, userId int, body string) error
	GetList(chatId int, userId int) ([]*models.Message, error)
}
