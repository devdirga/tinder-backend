package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "host=localhost user=postgres password=mysecretpassword dbname=tinder port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Database connection ping failed:", err)
	}
}
