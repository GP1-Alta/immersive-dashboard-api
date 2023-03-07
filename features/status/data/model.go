package data

import (
	_modelLog "immersive-dashboard/features/logs/data"
	_modelMentee "immersive-dashboard/features/mentees/data"

	"gorm.io/gorm"
)

// status
type Status struct {
	gorm.Model
	Name    string
	Mentees []_modelMentee.Mentee
	Logs    []_modelLog.Log
}
