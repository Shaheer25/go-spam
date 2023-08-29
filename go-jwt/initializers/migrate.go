package initializers

import "github.com/Shaheer25/go-jwt/model"

func MigrateFiles() {
	DB.AutoMigrate(&model.User{})
}