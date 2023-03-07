package data

import (
	"immersive-dashboard/features/logs"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	MenteeID   uint
	UserID     uint
	UserName   string
	StatusID   uint
	StatusName string
	Feedback   string
	CreatedAt  string
	Image      string
}

func CoreToLog(data logs.Core) Log {
	return Log{
		Model:      gorm.Model{ID: data.ID},
		MenteeID:   data.MenteeID,
		UserID:     data.UserID,
		UserName:   data.UserName,
		StatusID:   data.StatusID,
		StatusName: data.StatusName,
		Feedback:   data.Feedback,
		CreatedAt:  data.CreatedAt,
		Image:      data.Image,
	}
}

func LogToCore(data Log) logs.Core {
	return logs.Core{
		ID:         data.ID,
		MenteeID:   data.MenteeID,
		UserID:     data.UserID,
		UserName:   data.UserName,
		StatusID:   data.StatusID,
		StatusName: data.StatusName,
		Feedback:   data.Feedback,
		CreatedAt:  data.CreatedAt,
		Image:      data.Image,
	}
}

func ListLogToCore(data []Log) []logs.Core {
	var dataCore []logs.Core
	for _, v := range data {
		dataCore = append(dataCore, LogToCore(v))
	}
	return dataCore
}
