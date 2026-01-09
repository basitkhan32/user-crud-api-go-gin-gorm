package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(dbString string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
		return err
	}
	return nil
}

func DatabaseConnection() {
	var err error
	dbString := os.Getenv("DB_STRING")
	dbName := os.Getenv("DB_NAME")
	err = ConnectToDB(dbString)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "3D000" {
			_ = DB.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
			fmt.Println("Database created successfully")
			_ = ConnectToDB(dbString)
		}
	}

	log.Println("Database connected successfully")
}
