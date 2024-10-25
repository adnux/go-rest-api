package main

import (
	"fmt"
	"net/http"

	"github.com/adnux/go-rest-api/db"
	"github.com/adnux/go-rest-api/routes"
)

func main() {
	db.InitDB()
	server := http.NewServeMux()
	// server := gin.Default()

	routes.RegisterRoutes(server)

	// localhost:8080
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", server)

}
