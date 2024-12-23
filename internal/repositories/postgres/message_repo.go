package postgres

import (
	"boobook-chat-microservice/internal/models"
	"boobook-chat-microservice/internal/repositories"
	"database/sql"
	"fmt"
)

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repositories.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(chatId int, userId int, body string) error {
	const fnErr = "repository.postgres.messageRepository.Create"

	var err error

	// Проверить может ли пользователь просматривать сообщения в чате
	existsQuery := `SELECT EXISTS (SELECT 1 FROM chats_users WHERE chat_id = $1 AND user_id = $2)`
	var exists bool
	err = r.db.QueryRow(existsQuery, chatId, userId).Scan(&exists)
	if err != nil {
		return fmt.Errorf("%s: %w", fnErr, err)
	}
	if !exists {
		return repositories.ErrUserNotMember
	}

	stmt, err := r.db.Prepare("INSERT INTO messages (chat_id, user_id, body) VALUES ($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("(%s) error preparing insert statement: %w", fnErr, err)
	}

	_, err = stmt.Exec(chatId, userId, body)
	if err != nil {
		return fmt.Errorf("(%s) error executing user creation query: %w", fnErr, err)
	}

	return nil
}

func (r *messageRepository) GetList(chatId int, userId int) ([]*models.Message, error) {
	const fnErr = "repository.postgres.messageRepository.GetList"

	var err error
	var messages []*models.Message

	// Проверить может ли пользователь просматривать сообщения в чате
	existsQuery := `SELECT EXISTS (SELECT 1 FROM chats_users WHERE chat_id = $1 AND user_id = $2)`
	var exists bool
	err = r.db.QueryRow(existsQuery, chatId, userId).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}
	if !exists {
		return nil, repositories.ErrUserNotMember
	}

	query := `
		SELECT id, chat_id, user_id, body, created_at, updated_at 
		FROM messages
		WHERE chat_id = $1
	`

	var rows *sql.Rows

	rows, err = r.db.Query(query, chatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			err = fmt.Errorf("%s: %w", fnErr, err)
		}
	}()

	for rows.Next() {
		var message models.Message
		if err = rows.Scan(
			&message.ID,
			&message.ChatId,
			&message.UserId,
			&message.Body,
			&message.CreatedAt,
			&message.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", fnErr, err)
		}
		messages = append(messages, &message)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", fnErr, err)
	}

	return messages, nil
}
