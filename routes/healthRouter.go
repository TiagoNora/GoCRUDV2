package routes

import (
	"github.com/TiagoNora/GoCRUDV2/controller"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(engine *gin.Engine) {
	healthController := controller.NewHealthController()
	engine.GET("/health", healthController.Health)

}
