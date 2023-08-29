package main

import (
	"github.com/Shaheer_25/initializers"
	"github.com/Shaheer_25/models"
)

func init() {
	initializers.InitializeVar()
	initializers.ConnectToDB()

}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
