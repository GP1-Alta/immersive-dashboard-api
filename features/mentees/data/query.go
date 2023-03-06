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
	dataModel := CoreToModel(data)
	tx := repo.db.Where("id = ?", id).Delete(&dataModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete error, row affected = 0")
	}
	return nil
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
	var menteesModel []Mentee
	tx := repo.db.Find(&menteesModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	userCoreAll := ListModelToCore(menteesModel)
	return userCoreAll, nil
}

// Update implements mentees.MenteeDataInterface
func (repo *menteeQuery) Update(input mentees.Core, id uint) error {
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

func New(db *gorm.DB) mentees.MenteeDataInterface {
	return &menteeQuery{
		db: db,
	}
}
