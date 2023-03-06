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
	return ClassResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Mentor:       dataCore.Mentor,
		StartDateStr: dataCore.StartDateStr,
		EndDateStr:   dataCore.EndDateStr,
	}
}

func listCoreToResponse(dataCore []classes.Core) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
