package services

import (
	"boobook-chat-microservice/internal/http/requests"
	"boobook-chat-microservice/internal/models"
	"boobook-chat-microservice/internal/repositories"
	"fmt"
)

type messageService struct {
	messageRepository repositories.MessageRepository
}

func NewMessageService(messageRepository repositories.MessageRepository) MessageService {
	return &messageService{messageRepository: messageRepository}
}

func (s *messageService) Create(sendMessageRequest *requests.SendMessageRequest) error {
	const fnErr = "service.messageService.Create"

	err := s.messageRepository.Create(sendMessageRequest.ChatId, sendMessageRequest.UserId, sendMessageRequest.Body)
	if err != nil {
		return fmt.Errorf("%s: %w", fnErr, err)
	}

	return nil
}

func (s *messageService) GetList(chatId int, userId int) ([]*models.Message, error) {
	const fnErr = "service.messageService.GetList"

	messages, err := s.messageRepository.GetList(chatId, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return messages, nil
}
