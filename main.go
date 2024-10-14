package main

import (
	_ "embed"

	"github.com/adnux/go-rest-api/db"
	"github.com/adnux/go-rest-api/routes"
	"github.com/adnux/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.GetEnvironmentVariables()

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	// server.Use(cors.Default())

	server.Run(":8080") // localhost:8080
}
