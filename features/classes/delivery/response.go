package delivery

import "immersive-dashboard/features/classes"

type ClassResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Mentor       string `json:"mentor"`
	StartDateStr string `json:"start_date"`
	EndDateStr   string `json:"end_date"`
}

func coreToResponse(dataCore classes.Core) ClassResponse {
	startDate := dataCore.StartDate.Format("2006-01-02")
	endDate := dataCore.EndDate.Format("2006-01-02")
	return ClassResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Mentor:       dataCore.Mentor,
		StartDateStr: startDate,
		EndDateStr:   endDate,
	}
}

func listCoreToResponse(dataCore []classes.Core) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
