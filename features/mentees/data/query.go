package data

import (
	"errors"
	"fmt"
	"immersive-dashboard/features/mentees"

	"gorm.io/gorm"
)

type menteeQuery struct {
	db *gorm.DB
}

// Select implements mentees.MenteeDataInterface
func (repo *menteeQuery) Select(id uint) (mentees.Core, error) {
	var menteesModel Mentee
	tx := repo.db.Where("mentees.id = ?", id).Select("mentees.id, mentees.name, classes.name AS class, mentees.major, mentees.institution, mentees.phone, mentees.telegram, mentees.discord, mentees.email").Joins("JOIN classes ON classes.id = mentees.class_id").Find(&menteesModel)
	if tx.Error != nil {
		return mentees.Core{}, tx.Error
	}
	menteeCore := ModelToCore(menteesModel)
	return menteeCore, nil
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
func (repo *menteeQuery) SelectAll(limit, offset int, class, status, category, name string) ([]mentees.Core, error) {
	var whereClause string
	nameSearch := "%" + name + "%"
	classField := "classes.name"
	statusField := "statuses.name"
	categoryField := "mentees.category"

	//one filter
	if class != "" && status == "" && category == "" {
		whereClause = fmt.Sprintf("%s = '%s'", classField, class)
	} else if class == "" && status != "" && category == "" && name == "" {
		whereClause = fmt.Sprintf("%s = '%s'", statusField, status)
	} else if class == "" && status == "" && category != "" && name == "" {
		whereClause = fmt.Sprintf("%s = '%s'", categoryField, category)
	}

	//two filter
	if class != "" && status != "" && category == "" {
		whereClause = fmt.Sprintf("%s = '%s' AND %s = '%s'", classField, class, statusField, status)
	} else if class == "" && status != "" && category != "" {
		whereClause = fmt.Sprintf("%s = '%s' AND %s = '%s'", statusField, status, categoryField, category)
	} else if class != "" && status == "" && category != "" {
		whereClause = fmt.Sprintf("%s = '%s' AND %s = '%s'", classField, class, categoryField, category)
	}

	//three filter
	if class != "" && status != "" && category != "" {
		whereClause = fmt.Sprintf("%s = '%s' AND %s = '%s' AND %s = '%s'", classField, class, statusField, status, categoryField, category)
	}

	var menteesModel []Mentee
	if class == "" && status == "" && category == "" && name == "" {
		tx := repo.db.Limit(limit).Offset(offset).Order("mentees.name").Select("mentees.id, mentees.name, classes.name AS class, statuses.name AS status, mentees.category, mentees.sex").Joins("JOIN classes ON classes.id = mentees.class_id").Joins("JOIN statuses ON statuses.id = mentees.status_id").Find(&menteesModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if class == "" && status == "" && category == "" && name != "" {
		tx := repo.db.Limit(limit).Offset(offset).Where("mentees.name LIKE ?", nameSearch).Select("mentees.id, mentees.name, classes.name AS class, statuses.name AS status, mentees.category, mentees.sex").Joins("JOIN classes ON classes.id = mentees.class_id").Joins("JOIN statuses ON statuses.id = mentees.status_id").Find(&menteesModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := repo.db.Limit(limit).Offset(offset).Where(whereClause).Order("mentees.name").Select("mentees.id, mentees.name, classes.name AS class, statuses.name AS status, mentees.category, mentees.sex").Joins("JOIN classes ON classes.id = mentees.class_id").Joins("JOIN statuses ON statuses.id = mentees.status_id").Find(&menteesModel)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}
	menteeCoreAll := ListModelToCore(menteesModel)
	return menteeCoreAll, nil
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
