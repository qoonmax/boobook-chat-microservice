package main

import (
	"boobook-chat-microservice/internal/config"
	"boobook-chat-microservice/internal/repositories/postgres"
	"boobook-chat-microservice/internal/slogger"
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v4"
)

const (
	chatCount = 1000
)

func main() {
	logger := slogger.NewLogger()

	cfg := config.MustLoad()

	dbConnection, err := postgres.NewConnection(cfg.DBString)
	if err != nil {
		panic(fmt.Errorf("failed to connect to the database: %w", err))
	}

	defer func(dbConnection *sql.DB) {
		if err = postgres.CloseConnection(dbConnection); err != nil {
			logger.Error("failed to close the database connection", slogger.ErrorToSlogAttr(err))
			return
		}
	}(dbConnection)

	// Генерируем и вставляем чаты
	fmt.Println("Starting to insert chats...")

	chatsIds := make([]int, chatCount)

	for i := 0; i < chatCount; i++ {
		err = dbConnection.QueryRow("INSERT INTO chats DEFAULT VALUES RETURNING id").Scan(&chatsIds[i])
		if err != nil {
			logger.Error("failed to close the database connection", slogger.ErrorToSlogAttr(err))
			return
		}
	}

	fmt.Println("Chats have been successfully inserted.")

	// Свяжем чаты с пользователями
	fmt.Println("Starting to insert (chats_users) relations...")

	for i := 0; i < len(chatsIds); i++ {
		// Вставляем две записи для каждого чата
		_, err = dbConnection.Exec("INSERT INTO chats_users (chat_id, user_id) VALUES ($1, $2)", chatsIds[i], i+1)
		_, err = dbConnection.Exec("INSERT INTO chats_users (chat_id, user_id) VALUES ($1, $2)", chatsIds[i], i+2)
		if err != nil {
			logger.Error("failed to close the database connection", slogger.ErrorToSlogAttr(err))
			return
		}

		// Вставляем сообщения для каждого чата
		_, err = dbConnection.Exec("INSERT INTO messages (chat_id, user_id, body) VALUES ($1, $2, $3)", chatsIds[i], i+1, faker.Paragraph())
		_, err = dbConnection.Exec("INSERT INTO messages (chat_id, user_id, body) VALUES ($1, $2, $3)", chatsIds[i], i+2, faker.Paragraph())
		if err != nil {
			logger.Error("failed to close the database connection", slogger.ErrorToSlogAttr(err))
			return
		}
	}

	fmt.Println("Relations (chats_users) have been successfully inserted.")
}
