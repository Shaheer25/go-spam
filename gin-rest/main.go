package main

import (
	"github.com/Shaheer_25/controllers" // Corrected import path
	"github.com/Shaheer_25/initializers"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles"github.com/swaggo/files"
)

// @title My API
// @version 1.0
// @description This is a sample API with Swagger documentation

func init() {
	initializers.InitializeVar()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostShow)
	r.GET("/posts/:id", controllers.SinglePostShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	r.Run(":8080")
}
