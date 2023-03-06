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
	panic("unimplemented")
}

// Delete implements mentees.MenteeServiceInterface
func (service *menteeService) Delete(data mentees.Core, id uint) error {
	panic("unimplemented")
}

// Edit implements mentees.MenteeServiceInterface
func (service *menteeService) Edit(input mentees.Core, id uint) error {
	panic("unimplemented")
}

// GetAll implements mentees.MenteeServiceInterface
func (service *menteeService) GetAll() ([]mentees.Core, error) {
	panic("unimplemented")
}

func New(repo mentees.MenteeDataInterface) mentees.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
		validate:   validator.New(),
	}
}
