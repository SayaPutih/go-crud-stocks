package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "host=localhost port=5432 user=postgres password=kerlyn dbname=go_products sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB Open Error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	DB = db

	log.Println("Connected to PostgreSQL ✅")
}
