package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{} // In-memory storage for events

func (event Event) Save() {
	// todo: Implement logic to save the event to a database
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
