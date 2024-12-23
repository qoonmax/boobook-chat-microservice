package handlers

import (
	"github.com/gin-gonic/gin"
)

type MessageHandler interface {
	Create(ctx *gin.Context)
	GetList(ctx *gin.Context)
}
