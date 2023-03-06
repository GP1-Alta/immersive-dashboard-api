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
