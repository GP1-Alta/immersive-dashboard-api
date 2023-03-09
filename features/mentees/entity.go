package mentees

import "time"

type Core struct {
	ID              uint
	Name            string `validate:"required,max=50"`
	Address         string `validate:"required,max=50"`
	HomeAddress     string `validate:"required,max=50"`
	Email           string `validate:"required,email,max=50"`
	Sex             string `validate:"required"`
	Telegram        string `validate:"required,max=50"`
	Phone           string `validate:"required,max=12"`
	Discord         string `validate:"required,max=50"`
	StatusID        uint   `validate:"required"`
	Status          string
	ClassID         uint `validate:"required"`
	Class           string
	EmergencyName   string `validate:"required,max=50"`
	EmergencyPhone  string `validate:"required,max=12"`
	EmergencyStatus string `validate:"required,max=50"`
	Category        string `validate:"required,max=50"`
	Major           string `validate:"required,max=100"`
	Institution     string `validate:"required,max=100"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MenteeDataInterface interface {
	SelectAll(limit, offset int, class, status, category, name string) ([]Core, error)
	Select(id uint) (Core, error)
	Insert(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
}

type MenteeServiceInterface interface {
	GetAll(page int, class, status, category, name string) ([]Core, error)
	Get(id uint) (Core, error)
	Create(input Core) error
	Edit(input Core, id uint) error
	Delete(id uint) error
}
