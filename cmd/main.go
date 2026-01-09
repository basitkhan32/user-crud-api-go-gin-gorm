package main

import (
	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/routes"
	"github.com/basitkhan32/crud-api-go-gin/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// set gin mode
	gin.SetMode(gin.ReleaseMode)
	// setup env
	utils.LoadEnv()
	// setup database
	connections.DatabaseConnection()
	// setup router
	routes.Routes()
}
