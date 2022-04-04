package controller

import "github.com/gin-gonic/gin"

type IController interface {
	Inject(e *gin.Engine)
}
