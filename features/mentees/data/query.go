package data

import (
	"immersive-dashboard/features/mentees"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// Delete implements mentees.MenteeDataInterface
func (*menteeQuery) Delete(data mentees.Core, id uint) error {
	panic("unimplemented")
}

// Insert implements mentees.MenteeDataInterface
func (*menteeQuery) Insert(input mentees.Core) error {
	panic("unimplemented")
}

// SelectAll implements mentees.MenteeDataInterface
func (*menteeQuery) SelectAll() ([]mentees.Core, error) {
	panic("unimplemented")
}

// Update implements mentees.MenteeDataInterface
func (*menteeQuery) Update(input mentees.Core, id uint) error {
	panic("unimplemented")
}

func New(db *gorm.DB) mentees.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}
