package docs

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title My API
// @version 1.0
// @description This is a sample API with Swagger documentation.

// @host localhost:8080
// @BasePath /api/v1
func init() {
	r := gin.Default()

	// Swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
