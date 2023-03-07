package data

import (
	"immersive-dashboard/features/logs"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	MenteeID uint
	UserID   uint
	StatusID uint
	Feedback string
}

func CoreToLog(data logs.Core) Log {
	return Log{
		Model:    gorm.Model{ID: data.Id},
		MenteeID: data.MenteeID,
		UserID:   data.UserID,
		StatusID: data.StatusID,
		Feedback: data.Feedback,
	}
}

func LogToCore(data Log) logs.Core {
	return logs.Core{
		Id:       data.ID,
		MenteeID: data.MenteeID,
		UserID:   data.UserID,
		StatusID: data.StatusID,
		Feedback: data.Feedback,
	}
}

func ListLogToCore(data []Log) []logs.Core {
	var dataCore []logs.Core
	for _, v := range data {
		dataCore = append(dataCore, LogToCore(v))
	}
	return dataCore
}
