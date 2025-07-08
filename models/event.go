package models

import "time"

type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserID      int       `json:"user_id"`
}

var events = []Event{} // In-memory storage for events

func (event Event) Save() {
	// todo: Implement logic to save the event to a database
	events = append(events, event)
}
