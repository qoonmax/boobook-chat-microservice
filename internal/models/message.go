package models

import "time"

type Message struct {
	ID        int       `json:"id"`
	ChatId    int       `json:"chat_id"`
	UserId    int       `json:"user_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
