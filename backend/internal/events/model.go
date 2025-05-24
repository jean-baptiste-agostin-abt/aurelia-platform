package events

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	Type      string
	CapsuleID uint
	UserID    uint
}
