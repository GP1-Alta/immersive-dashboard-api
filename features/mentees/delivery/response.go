package delivery

import "immersive-dashboard/features/mentees"

type MenteeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Status   string `json:"status"`
	Class    string `json:"class"`
	Category string `json:"category"`
}

func coreToResponse(dataCore mentees.Core) MenteeResponse {
	return MenteeResponse{
		ID:       dataCore.ID,
		Name:     dataCore.Name,
		Sex:      dataCore.Sex,
		Status:   dataCore.Status,
		Class:    dataCore.Class,
		Category: dataCore.Category,
	}
}

func listCoreToResponse(dataCore []mentees.Core) []MenteeResponse {
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
