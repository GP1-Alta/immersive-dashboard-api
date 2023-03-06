package data

import (
	"errors"
	"immersive-dashboard/features/mentees"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// Delete implements mentees.MenteeDataInterface
func (repo *menteeQuery) Delete(data mentees.Core, id uint) error {
	panic("unimplemented")
}

// Insert implements mentees.MenteeDataInterface
func (repo *menteeQuery) Insert(input mentees.Core) error {
	dataModel := CoreToModel(input)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert error, row affected = 0")
	}
	return nil
}

// SelectAll implements mentees.MenteeDataInterface
func (repo *menteeQuery) SelectAll() ([]mentees.Core, error) {
	panic("unimplemented")
}

// Update implements mentees.MenteeDataInterface
func (repo *menteeQuery) Update(input mentees.Core, id uint) error {
	panic("unimplemented")
}

func New(db *gorm.DB) mentees.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}
