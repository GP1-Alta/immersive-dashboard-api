package delivery

import "immersive-dashboard/features/mentees"

type MenteeRequest struct {
	Name            string `json:"name" form:"name"`
	Address         string `json:"address" form:"address"`
	HomeAddress     string `json:"home_address" form:"home_address"`
	Email           string `json:"email" form:"email"`
	Sex             string `json:"sex" form:"sex"`
	Telegram        string `json:"telegram" form:"telegram"`
	Phone           string `json:"phone" form:"phone"`
	Discord         string `json:"discord" form:"discord"`
	StatusID        uint   `json:"status" form:"status"`
	ClassID         uint   `json:"class" from:"class"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Category        string `json:"category" form:"category"`
	Major           string `json:"major" form:"major"`
	Institution     string `json:"institution" form:"institution"`
}

func requestToCore(dataRequest MenteeRequest) mentees.Core {
	return mentees.Core{
		Name:            dataRequest.Name,
		Address:         dataRequest.Address,
		HomeAddress:     dataRequest.HomeAddress,
		Email:           dataRequest.Email,
		Sex:             dataRequest.Sex,
		Telegram:        dataRequest.Telegram,
		Phone:           dataRequest.Phone,
		Discord:         dataRequest.Discord,
		StatusID:        dataRequest.StatusID,
		ClassID:         dataRequest.ClassID,
		EmergencyName:   dataRequest.EmergencyName,
		EmergencyPhone:  dataRequest.EmergencyPhone,
		EmergencyStatus: dataRequest.EmergencyStatus,
		Category:        dataRequest.Category,
		Major:           dataRequest.Major,
		Institution:     dataRequest.Institution,
	}
}
