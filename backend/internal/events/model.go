package events

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Type      string
	CapsuleID uint
	UserID    uint
}

func NewEvent(eventType string, capsuleID, userID uint) *Event {
	return &Event{Type: eventType, CapsuleID: capsuleID, UserID: userID}
}
