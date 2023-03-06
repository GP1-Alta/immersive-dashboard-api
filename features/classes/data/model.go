package data

import (
	"immersive-dashboard/features/classes"
	_modelMentee "immersive-dashboard/features/mentees/data"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name         string
	UserID       uint
	Mentor       string `gorm:"-:migration"`
	StartDate    time.Time
	StartDateStr string `gorm:"-:migration"`
	EndDate      time.Time
	EndDateStr   string `gorm:"-:migration"`
	Mentees      []_modelMentee.Mentee
}

func CoreToModel(dataCore classes.Core) Class {
	return Class{
		Name:         dataCore.Name,
		UserID:       dataCore.UserID,
		Mentor:       dataCore.Mentor,
		StartDate:    dataCore.StartDate,
		StartDateStr: dataCore.StartDateStr,
		EndDate:      dataCore.EndDate,
		EndDateStr:   dataCore.EndDateStr,
	}
}

func ModelToCore(dataModel Class) classes.Core {
	return classes.Core{
		ID:           dataModel.ID,
		Name:         dataModel.Name,
		UserID:       dataModel.UserID,
		Mentor:       dataModel.Mentor,
		StartDate:    dataModel.StartDate,
		StartDateStr: dataModel.StartDateStr,
		EndDate:      dataModel.EndDate,
		EndDateStr:   dataModel.EndDateStr,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
	}
}

func ListModelToCore(dataModel []Class) []classes.Core {
	var dataCore []classes.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
