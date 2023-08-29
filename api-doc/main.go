package main

import (
	

	_ "github.com/Shaheer25/api-doc/docs"
	"github.com/Shaheer25/api-doc/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API
// @version 1
// @Description Sample description

// @contact.name Shaheer
// @contact.url https://github.com/Shaheer25
// @contact.email gmail@gmail.com

// @host localhost:8080
// @BasePath /api/v1

func main() {

	r := gin.Default()

	user := r.Group("api/v1/users")
	{

		user.GET("/", handlers.GetUsers)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	 r.Run()
	
}
