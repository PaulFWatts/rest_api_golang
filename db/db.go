package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Importing SQLite driver
)

var DB *sql.DB // Global variable to hold the database connection

// InitDB initializes the database connection
// This function should be called at the start of the application
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10) // Set maximum open connections to the database
	DB.SetMaxIdleConns(5)  // Set maximum idle connections to the database
	createTables()         // Create necessary tables if they do not exist
}

// CreateTables creates necessary tables in the database
func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
		)
		`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Failed to create events table: " + err.Error())
	}

}
