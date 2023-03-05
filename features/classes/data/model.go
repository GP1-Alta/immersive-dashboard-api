package data

import (
	_modelMentee "immersive-dashboard/features/mentees/data"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string
	UserID    uint
	StartDate time.Time
	EndDate   time.Time
	Mentees   []_modelMentee.Mentee
}
