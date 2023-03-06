package data

import (
	"errors"
	"immersive-dashboard/features/classes"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Delete implements classes.ClassDataInterface
func (repo *classQuery) Delete(data classes.Core, id uint) error {
	panic("unimplemented")
}

// Insert implements classes.ClassDataInterface
func (repo *classQuery) Insert(input classes.Core) error {
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

// ListAll implements classes.ClassDataInterface
func (repo *classQuery) ListAll() ([]classes.Core, error) {
	panic("unimplemented")
}

// SelectAll implements classes.ClassDataInterface
func (repo *classQuery) SelectAll(limit int, offset int, name string) ([]classes.Core, error) {
	panic("unimplemented")
}

// Update implements classes.ClassDataInterface
func (repo *classQuery) Update(input classes.Core, id uint) error {
	panic("unimplemented")
}

func New(db *gorm.DB) classes.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
