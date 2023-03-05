package data

import "gorm.io/gorm"

type Mentee struct {
	gorm.Model
	Email    string
	Name     string
	Sex      string
	Log      uint
	Status   string `gorm:"type:enum('Interview', 'Join Class', 'Unit 1', 'Unit 2', 'Unit 3', 'Repeat Unit 1', 'Repeat Unit 2', 'Repeat Unit 3', 'Placement', 'Eliminated', 'Graduated')"`
	Class    uint
	Category string `gorm:"type:enum('Informatics','Non-Informatics')"`
}
