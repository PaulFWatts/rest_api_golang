package models

import (
	"github.com/PaulFWatts/rest_api_golang/db"
	"time"
)


type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{} // In-memory storage for events

func (e Event) Save() error {
	query := `
	INSERT INTO events (id, name, description, location, datetime, user_id) 
	VALUES (?, ?, ?, ?, ?, ?)` // SQL query to insert a new event
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err // Return error if the query fails
	}
	defer stmt.Close() // Ensure the statement is closed after execution
	result, err := stmt.Exec(e.ID, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err // Return error if execution fails
	}
	id, err := result.LastInsertId() // Get the last inserted ID
	e.ID = id
	return err // Return nil if everything is successful
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events" // SQL query to retrieve all events
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err // Return err if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after reading
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err // Return err if scanning fails
		}
		events = append(events, event) // Append the event to the slice
	}
	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	// SQL query to retrieve a specific event by ID
	
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) // Query for a specific event by ID
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err // Return an empty event and error if not found
	}
	return &event, nil // Return the found event
}

func (event Event) Update() error {
	// SQL query to update an existing event
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err // Return error if the query preparation fails
	}
	defer stmt.Close() // Ensure the statement is closed after execution
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err // Return nil if everything is successful
}