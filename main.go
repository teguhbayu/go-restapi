package main

import (
	"github.com/teguhbayu/go-restapi/initializers"
	"github.com/teguhbayu/go-restapi/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.InitDb()
}

func main() {
	routes.NewRoutes().Run()
}
