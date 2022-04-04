package main

import (
	"gin-app/controller"
	"gin-app/service"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	basicService := service.NewBasicService("Hello from DI")
	bc := controller.NewBasicController(basicService)
	bc.Inject(engine)
	engine.Run()
}
