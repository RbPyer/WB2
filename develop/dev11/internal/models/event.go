package models

import "time"

func NewEvent(uuid string, date time.Time) *Event {
	return &Event{
		UUID: uuid,
		Date: date,
	}
}

type Event struct {
	UUID string    `json:"event_id"`
	Date time.Time `json:"date"`
}
