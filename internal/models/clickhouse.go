package models

import "time"

type Events struct {
	EventID   int
	EventType string
}

type EventHandler struct {
	EventType string `json:"eventType"`
	UserID    int    `json:"userID"`
	EventTime string `json:"eventTime"`
	Payload   string `json:"payload"`
}

type EventRepository struct {
	EventID   int
	EventType string
	UserID    int
	EventTime time.Time
	Payload   string
}
