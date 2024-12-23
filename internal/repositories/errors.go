package repositories

import "errors"

var (
	ErrUserNotMember = errors.New("user is not a member of the chat")
)
