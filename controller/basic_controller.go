package controller

import (
	"gin-app/service"

	"github.com/gin-gonic/gin"
)

type BasicController struct {
	IBasicService service.IBasicService
}

func NewBasicController(bs service.IBasicService) IController {
	return BasicController{IBasicService: bs}
}

func (bc BasicController) Inject(e *gin.Engine) {
	e.GET("/basic", bc.IBasicService.GetBasicResponse)
}
