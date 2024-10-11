package main

import (
	"github.com/adnux/go-rest-api/db"
	"github.com/adnux/go-rest-api/routes"
	schemas "github.com/adnux/go-rest-api/schema"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	schemas.CreateTables()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
