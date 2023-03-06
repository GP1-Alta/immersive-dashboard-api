package service

import (
	"immersive-dashboard/features/mentees"

	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeData mentees.MenteeDataInterface
	validate   *validator.Validate
}

// Create implements mentees.MenteeServiceInterface
func (service *menteeService) Create(input mentees.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.menteeData.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

// Delete implements mentees.MenteeServiceInterface
func (service *menteeService) Delete(data mentees.Core, id uint) error {
	errDelete := service.menteeData.Delete(data, id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

// Edit implements mentees.MenteeServiceInterface
func (service *menteeService) Edit(input mentees.Core, id uint) error {
	errUpdate := service.menteeData.Update(input, id)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}

// GetAll implements mentees.MenteeServiceInterface
func (service *menteeService) GetAll(page int) ([]mentees.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := service.menteeData.SelectAll(limit, offset)
	return data, err
}

func New(repo mentees.MenteeDataInterface) mentees.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
		validate:   validator.New(),
	}
}
