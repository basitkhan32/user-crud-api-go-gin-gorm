package db

import (
	"log"
)

// DBMigration runs migrations for the provided models
// This function accepts models as interface{} to avoid import cycles
func DBMigration(models ...interface{}) {
	if len(models) == 0 {
		log.Println("No models provided for migration")
		return
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration done successfully")
}
