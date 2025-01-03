package db

import (
	"database/sql"
	"log"
	"os"
	"github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB will initialize the SQLite database connection
func InitDB() {
	var err error

	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "./tasks.db" // default to local tasks.db file
	}

	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("Error opening SQLite databse: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL DEFAULT FLASE
	); 
	`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating tasks table: %v", err)
	}

	log.Println("Successfully connected to the SQLite database.")
}

func GetDB() *sql.DB {
	return db
}
