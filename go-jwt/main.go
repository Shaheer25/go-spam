package main

import (
	controllers "github.com/Shaheer25/go-jwt/controller"
	"github.com/Shaheer25/go-jwt/initializers"
	"github.com/Shaheer25/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateFiles()
}

func main() {
	r := gin.Default()
	r.POST("/Signup", controllers.Signup)
	r.POST("/Login", controllers.Login)
	r.GET("/Validate", middleware.RequiredAuth, controllers.Validate)
	r.Run()
}
