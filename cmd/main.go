package main

import (
	"fmt"
	"os"

	"github.com/basitkhan32/crud-api-go-gin/db"
	"github.com/basitkhan32/crud-api-go-gin/routes"
	"github.com/basitkhan32/crud-api-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// set gin mode
	gin.SetMode(gin.ReleaseMode)
	// setup env
	utils.LoadEnv()
	// setup database
	db.DatabaseConnection()

	// apply migration
	// migrations.DBMigration()

	// setup router
	routes.Routes(r)

	PORT := ""
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Server is running on port " + PORT)
	r.Run(":" + PORT)
}
