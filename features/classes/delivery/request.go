package delivery

import (
	"immersive-dashboard/features/classes"
	"time"
)

type ClassRequest struct {
	Name         string `json:"name" form:"name"`
	UserID       uint   `json:"user_id" form:"user_id"`
	StartDateStr string `json:"start_date" form:"start_date"`
	EndDateStr   string `json:"end_date" form:"end_date"`
}

func requestToCore(dataRequest ClassRequest) classes.Core {
	startDate, _ := time.Parse("2006-01-02", dataRequest.StartDateStr)
	endDate, _ := time.Parse("2006-01-02", dataRequest.EndDateStr)
	return classes.Core{
		Name:      dataRequest.Name,
		UserID:    dataRequest.UserID,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
