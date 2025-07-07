package main

import (
	"github.com/teguhbayu/go-restapi/initializers"
	"github.com/teguhbayu/go-restapi/models"
)

func init() {
	initializers.LoadEnv()
	initializers.InitDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Books{})
}
