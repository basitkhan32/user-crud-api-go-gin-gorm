package routes

import (
	"log"
	"os"

	"github.com/basitkhan32/crud-api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.POST("/create", controllers.CreateUserController)
	r.GET("/read", controllers.ReadUserController)
	r.GET("/read/:id", controllers.ReadUserIDController)
	r.PUT("/update/:id ", controllers.UpdateUserController)
	r.DELETE("/delete/:id", controllers.DeleteUserController)
	port := os.Getenv("PORT")
	log.Println("Router is working at port:", port)
	r.Run(":" + port)
}
