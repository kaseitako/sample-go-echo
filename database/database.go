package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./sample.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Successfully connected to SQLite database")
	runMigrations()
}

func runMigrations() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(createUserTable); err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	log.Println("Database migrations completed successfully")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}