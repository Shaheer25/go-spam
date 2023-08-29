package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	credentials := os.Getenv("CREDENTIALS")
	DB, err = gorm.Open(postgres.Open(credentials), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connecting to DataBase")
	}

}
