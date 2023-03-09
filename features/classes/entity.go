package classes

import "time"

type Core struct {
	ID           uint
	Name         string `validate:"required,max=50"`
	UserID       uint   `validate:"required"`
	Mentor       string
	StartDate    time.Time `validate:"required"`
	StartDateStr string
	EndDate      time.Time `validate:"required"`
	EndDateStr   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ClassDataInterface interface {
	SelectAll(limit, offset int, name string) ([]Core, error)
	SelectOne(id uint) (Core, error)
	ListAll() ([]Core, error)
	Insert(input Core) error
	Update(input Core, id uint) error
	Delete(id uint) error
}

type ClassServiceInterface interface {
	GetAll(page int, name string) ([]Core, error)
	GetOne(id uint) (Core, error)
	List() ([]Core, error)
	Create(input Core) error
	Edit(input Core, id uint) error
	Delete(id uint) error
}
