package data

import (
	_modelLog "immersive-dashboard/features/logs/data"

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
	ClassID         uint
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string `gorm:"type:enum('Orang Tua', 'Kakek Nenek', 'Saudara dari Orang Tua')"`
	Category        string `gorm:"type:enum('Informatics','Non-Informatics')"`
	Major           string
	Institution     string
	Logs            []_modelLog.Log
}
