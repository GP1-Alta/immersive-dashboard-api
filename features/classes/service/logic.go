package service

import (
	"immersive-dashboard/features/classes"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData classes.ClassDataInterface
	validate  *validator.Validate
}

// GetOne implements classes.ClassServiceInterface
func (service *classService) GetOne(id uint) (classes.Core, error) {
	data, err := service.classData.SelectOne(id)
	return data, err
}

// Create implements classes.ClassServiceInterface
func (service *classService) Create(input classes.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.classData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// Delete implements classes.ClassServiceInterface
func (service *classService) Delete(data classes.Core, id uint) error {
	errDelete := service.classData.Delete(data, id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

// Edit implements classes.ClassServiceInterface
func (service *classService) Edit(input classes.Core, id uint) error {
	errUpdate := service.classData.Update(input, id)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

// GetAll implements classes.ClassServiceInterface
func (service *classService) GetAll(page int, name string) ([]classes.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := service.classData.SelectAll(limit, offset, name)
	return data, err
}

// List implements classes.ClassServiceInterface
func (service *classService) List() ([]classes.Core, error) {
	data, err := service.classData.ListAll()
	return data, err
}

func New(repo classes.ClassDataInterface) classes.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}
