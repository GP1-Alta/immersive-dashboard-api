package data

import (
	_modelMentee "immersive-dashboard/features/mentees/data"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Status   string `gorm:"type:enum('Interview', 'Join Class', 'Unit 1', 'Unit 2', 'Unit 3', 'Repeat Unit 1', 'Repeat Unit 2', 'Repeat Unit 3', 'Placement', 'Eliminated', 'Graduated')"`
	Feedback string
	Mentees  []_modelMentee.Mentee `gorm:"foreignKey:Log"`
}
