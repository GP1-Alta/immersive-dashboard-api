package data

import (
	_modelLog "immersive-dashboard/features/logs/data"
	"immersive-dashboard/features/mentees"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	Name            string
	Address         string
	HomeAddress     string
	Email           string
	Sex             string `gorm:"type:enum('Male', 'Female')"`
	Telegram        string
	Phone           string
	Discord         string
	StatusID        uint
	Status          string `gorm:"-:migration"`
	ClassID         uint
	Class           string `gorm:"-:migration"`
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string `gorm:"type:enum('Orang Tua', 'Kakek Nenek', 'Saudara dari Orang Tua')"`
	Category        string `gorm:"type:enum('Informatics','Non-Informatics')"`
	Major           string
	Institution     string
	Logs            []_modelLog.Log
}

func CoreToModel(dataCore mentees.Core) Mentee {
	return Mentee{
		Name:            dataCore.Name,
		Address:         dataCore.Address,
		HomeAddress:     dataCore.HomeAddress,
		Email:           dataCore.Email,
		Sex:             dataCore.Sex,
		Telegram:        dataCore.Telegram,
		Phone:           dataCore.Phone,
		Discord:         dataCore.Discord,
		StatusID:        dataCore.StatusID,
		Status:          dataCore.Status,
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

func ModelToCore(dataModel Mentee) mentees.Core {
	return mentees.Core{
		ID:              dataModel.ID,
		Name:            dataModel.Name,
		Address:         dataModel.Address,
		HomeAddress:     dataModel.HomeAddress,
		Email:           dataModel.Email,
		Sex:             dataModel.Sex,
		Telegram:        dataModel.Telegram,
		Phone:           dataModel.Phone,
		Discord:         dataModel.Discord,
		StatusID:        dataModel.StatusID,
		Status:          dataModel.Status,
		ClassID:         dataModel.ClassID,
		Class:           dataModel.Class,
		EmergencyName:   dataModel.EmergencyName,
		EmergencyPhone:  dataModel.EmergencyPhone,
		EmergencyStatus: dataModel.EmergencyStatus,
		Category:        dataModel.Category,
		Major:           dataModel.Major,
		Institution:     dataModel.Institution,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}

func ListModelToCore(dataModel []Mentee) []mentees.Core {
	var dataCore []mentees.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, ModelToCore(v))
	}
	return dataCore
}
