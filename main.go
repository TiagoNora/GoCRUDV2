package main

import (
	"github.com/TiagoNora/GoCRUDV2/config/db"
	"github.com/TiagoNora/GoCRUDV2/config/logger"
	"github.com/TiagoNora/GoCRUDV2/config/minioClient"
	"github.com/TiagoNora/GoCRUDV2/consumers"
	"github.com/TiagoNora/GoCRUDV2/docs"
	routes "github.com/TiagoNora/GoCRUDV2/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title           Swagger GO API
// @version         1.0
// @description     This is a example GO API.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

func init() {
	logger.InitLogger()
	db.ConnectDatabase()
	consumers.ConsumeTopics()
	minioClient.NewMinioClient()
}

func main() {
	basePath := "/"
	docs.SwaggerInfo.BasePath = basePath
	err := godotenv.Load()
	if err != nil {
		return
	}
	port := os.Getenv("PORT")

	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.AllRoutes(router)

	err = router.Run(":" + port)
	if err != nil {
		return
	}
}
