package data

import (
	_modelClass "immersive-dashboard/features/classes/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string `gorm:"type:enum('Admission', 'Mentor', 'People Skill', 'Placement', 'Academic')"`
	Team     string `gorm:"type:enum('Active', 'Not-Active', 'Deleted')"`
	Status   string
	Classes  []_modelClass.Class `gorm:"foreignKey:Mentor"`
}
