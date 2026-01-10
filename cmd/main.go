package main

import (
	"fmt"
	"log"
	"os"

	"github.com/basitkhan32/crud-api-go-gin/db"
	"github.com/basitkhan32/crud-api-go-gin/modules/user"
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
	err := db.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration done successfully")

	// setup router
	routes.Routes(r)

	PORT := ""
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Server is running on port " + PORT)
	r.Run(":" + PORT)
}
