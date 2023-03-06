package mentees

import "time"

type Core struct {
	ID        uint
	Email     string `validate:"required,email"`
	Name      string `validate:"required"`
	Sex       string `validate:"required"`
	Status    string `validate:"required"`
	ClassID   string `validate:"required"`
	Category  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
