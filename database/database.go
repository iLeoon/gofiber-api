package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = ""
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = ""
)

func Connect() (*sql.DB, error) {
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
