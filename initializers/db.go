package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error

	dsn := os.Getenv("DB_HOST")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("db err!")

	}

}
