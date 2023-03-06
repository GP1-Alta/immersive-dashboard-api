package service

import (
	"immersive-dashboard/features/classes"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData classes.ClassDataInterface
	validate  *validator.Validate
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
	panic("unimplemented")
}

// Edit implements classes.ClassServiceInterface
func (service *classService) Edit(input classes.Core, id uint) error {
	panic("unimplemented")
}

// GetAll implements classes.ClassServiceInterface
func (service *classService) GetAll(page int, name string) ([]classes.Core, error) {
	panic("unimplemented")
}

// List implements classes.ClassServiceInterface
func (service *classService) List() ([]classes.Core, error) {
	panic("unimplemented")
}

func New(repo classes.ClassDataInterface) classes.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}
