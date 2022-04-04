package service

import (
	"github.com/gin-gonic/gin"
)

type BasicService struct {
	Message string
}

type IBasicService interface {
	GetBasicResponse(context *gin.Context)
}

func NewBasicService(message string) IBasicService {
	return &BasicService{Message: message}
}

func (i *BasicService) GetBasicResponse(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": i.Message,
	})
}
