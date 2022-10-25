package main

import (
	"github.com/COOPSPROFI/GoLangProject/pkg/configs"
	"github.com/COOPSPROFI/GoLangProject/pkg/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	// Connect all configs
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	configs.SyncDatabase()
}

func main() {
	// init Router
	r := gin.Default()

	// Image url browse
	r.Static("/assets", "./assets")

	// All routes of application
	routes.RegisterEventRoutes(r)
	routes.RegisterPrinterRoutes(r)
	routes.RegisterUserRoutes(r)

	// Run the Server
	r.Run()
}
