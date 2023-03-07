package data

import (
	_modelClass "immersive-dashboard/features/classes/data"
	_modelLog "immersive-dashboard/features/logs/data"
	"immersive-dashboard/features/users"

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

func CoreToUser(data users.Core) User {
	return User{
		Model:    gorm.Model{ID: data.Id},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Status:   data.Status,
		Role:     data.Role,
	}
}

func UserToCore(data User) users.Core {
	return users.Core{
		Id:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Status:   data.Status,
		Role:     data.Role,
	}
}

func ListUserToCore(data []User) []users.Core {
	var dataCore []users.Core
	for _, v := range data {
		dataCore = append(dataCore, UserToCore(v))
	}
	return dataCore
}