package data

import (
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	MenteeID uint
	UserID   uint
	StatusID uint
	Feedback string
}
