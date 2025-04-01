package routes

import (
	"github.com/TiagoNora/GoCRUDV2/controller"
	"github.com/gin-gonic/gin"
)

func ImageRoutes(engine *gin.Engine) {
	imageController := controller.NewImageController()
	engine.GET("/image/:fileName", imageController.GetImage)
	engine.POST("/image", imageController.CreateImage)
}
