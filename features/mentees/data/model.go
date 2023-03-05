package data

import (
	_modelLog "immersive-dashboard/features/logs/data"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Email    string
	Name     string
	Sex      string
	Status   string `gorm:"type:enum('Interview', 'Join Class', 'Unit 1', 'Unit 2', 'Unit 3', 'Repeat Unit 1', 'Repeat Unit 2', 'Repeat Unit 3', 'Placement', 'Eliminated', 'Graduated')"`
	ClassID  uint
	Category string `gorm:"type:enum('Informatics','Non-Informatics')"`
	Logs     []_modelLog.Log
}
