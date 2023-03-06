package data

import (
	_modelClass "immersive-dashboard/features/classes/data"
	_modelLog "immersive-dashboard/features/logs/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Team     string `gorm:"type:enum('Admission', 'Mentor', 'People Skill', 'Placement', 'Academic')"`
	Status   string `gorm:"type:enum('Active', 'Not-Active', 'Deleted')"`
	Role     string `gorm:"type:varchar(10) not null default 'Default'"`
	Classes  []_modelClass.Class
	Logs     []_modelLog.Log
}
