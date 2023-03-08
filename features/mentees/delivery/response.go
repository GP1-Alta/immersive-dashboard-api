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

type DetailMenteeResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	HomeAddress     string `json:"home_address"`
	Email           string `json:"email"`
	Sex             string `json:"sex"`
	Telegram        string `json:"telegram"`
	Phone           string `json:"phone"`
	Discord         string `json:"discord"`
	StatusID        uint   `json:"status_id"`
	ClassID         uint   `json:"class_id"`
	Class           string `json:"class"`
	EmergencyName   string `json:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status"`
	Category        string `json:"category"`
	Major           string `json:"major"`
	Institution     string `json:"institution"`
}

func coreToDetailMenteeResponse(dataCore mentees.Core) DetailMenteeResponse {
	return DetailMenteeResponse{
		ID:              dataCore.ID,
		Name:            dataCore.Name,
		Address:         dataCore.Address,
		HomeAddress:     dataCore.HomeAddress,
		Email:           dataCore.Email,
		Sex:             dataCore.Sex,
		Telegram:        dataCore.Telegram,
		Phone:           dataCore.Phone,
		Discord:         dataCore.Discord,
		StatusID:        dataCore.StatusID,
		ClassID:         dataCore.ClassID,
		Class:           dataCore.Class,
		EmergencyName:   dataCore.EmergencyName,
		EmergencyPhone:  dataCore.EmergencyPhone,
		EmergencyStatus: dataCore.EmergencyStatus,
		Category:        dataCore.Category,
		Major:           dataCore.Major,
		Institution:     dataCore.Institution,
	}
}
