package routes

import (
	"github.com/TiagoNora/GoCRUDV2/controller"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(engine *gin.Engine) {
	productController := controller.NewProductController()

	engine.GET("/product/:id", productController.FindById)
	engine.POST("/product", productController.Create)
	engine.PUT("/product/:id", productController.Update)
	engine.DELETE("/product/:id", productController.Delete)
	engine.GET("/products", productController.FindAll)
}
