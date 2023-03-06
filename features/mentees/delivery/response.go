package delivery

import "immersive-dashboard/features/mentees"

type MenteeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	StatusID uint   `json:"status_id"`
	ClassID  uint   `json:"class_id"`
	Category string `json:"category"`
}

func coreToResponse(dataCore mentees.Core) MenteeResponse {
	return MenteeResponse{
		ID:       dataCore.ID,
		Name:     dataCore.Name,
		Sex:      dataCore.Sex,
		StatusID: dataCore.StatusID,
		ClassID:  dataCore.ClassID,
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
