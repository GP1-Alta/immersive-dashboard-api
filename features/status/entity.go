package status

import "time"

type Core struct {
	ID        uint
	Name      string `validate:"required,max=50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
