package requests

type SendMessageRequest struct {
	ChatId int    `json:"chat_id"`
	UserId int    `json:"user_id"`
	Body   string `json:"body" binding:"required,max=512"`
}
