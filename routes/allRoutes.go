package routes

import (
	"github.com/gin-gonic/gin"
)

func AllRoutes(engine *gin.Engine) {
	ProductRoutes(engine)
	HealthRoutes(engine)
	ImageRoutes(engine)
}
