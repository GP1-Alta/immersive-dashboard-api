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
	var classesModel []Class
	tx := repo.db.Where("classes.name LIKE ?", name).Select("classes.id, classes.name, users.name AS mentor, classes.start_date, classes.end_date").Joins("JOIN users ON classes.user_id = users.id").Find(&classesModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	classCoreAll := ListModelToCore(classesModel)
	return classCoreAll, nil
}

// Update implements classes.ClassDataInterface
func (repo *classQuery) Update(input classes.Core, id uint) error {
	dataModel := CoreToModel(input)
	tx := repo.db.Where("id = ?", id).Updates(&dataModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update error, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) classes.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
