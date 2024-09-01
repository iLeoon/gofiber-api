package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Can't reload env variables")
	}

	host := os.Getenv("DATABASE_HOST")
	port := 5432
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Something went wrong while tring to attempt the connection: %w", err)
	}

	fmt.Println("Connected successfully to the database")
	return db, nil
}
