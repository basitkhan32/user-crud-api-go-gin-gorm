package db

import (
	"log"

	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/modules/user"
)

func DBMigration() {
	// Auto migrate the User model
	err := connections.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration done successfully")
}
