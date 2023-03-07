package delivery

import "immersive-dashboard/features/status"

type StatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func coreToResponse(dataCore status.Core) StatusResponse {
	return StatusResponse{
		ID:   dataCore.ID,
		Name: dataCore.Name,
	}
}

func listCoreToResponse(dataCore []status.Core) []StatusResponse {
	var dataResponse []StatusResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, coreToResponse(v))
	}
	return dataResponse
}
