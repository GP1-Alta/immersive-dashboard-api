package mentees

import "time"

type Core struct {
	ID              uint
	Name            string
	Address         string
	HomeAddress     string
	Email           string
	Sex             string
	Telegram        string
	Phone           string
	Discord         string
	StatusID        uint
	ClassID         uint
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	Category        string
	Major           string
	Institution     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MenteeDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	Update(input Core, id uint) error
	Delete(data Core, id uint) error
}

type MenteeServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
	Edit(input Core, id uint) error
	Delete(data Core, id uint) error
}
