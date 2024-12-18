package model

import (
	"database/sql"
	"gotinder/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", config.GetConf().DB)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Database connection ping failed:", err)
	}
}
