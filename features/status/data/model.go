package data

import (
	_modelLog "immersive-dashboard/features/logs/data"
	_modelMentee "immersive-dashboard/features/mentees/data"
	"immersive-dashboard/features/status"

	"gorm.io/gorm"
)

// status
type Status struct {
	gorm.Model
	Name    string
	Mentees []_modelMentee.Mentee
	Logs    []_modelLog.Log
}

func CoreToModel(dataCore status.Core) Status {
	return Status{
		Name: dataCore.Name,
	}
}

func ModelToCore(dataModel Status) status.Core {
	return status.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func ListModelToCore(dataModel []Status) []status.Core {
	var dataCore []status.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
