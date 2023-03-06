package service

import (
	"immersive-dashboard/features/classes"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData classes.ClassDataInterface
	validate  *validator.Validate
}

// Delete implements classes.ClassDataInterface
func (*classService) Delete(data classes.Core, id uint) error {
	panic("unimplemented")
}

// Insert implements classes.ClassDataInterface
func (*classService) Insert(input classes.Core) error {
	panic("unimplemented")
}

// ListAll implements classes.ClassDataInterface
func (*classService) ListAll() ([]classes.Core, error) {
	panic("unimplemented")
}

// SelectAll implements classes.ClassDataInterface
func (*classService) SelectAll(limit int, offset int, name string) ([]classes.Core, error) {
	panic("unimplemented")
}

// Update implements classes.ClassDataInterface
func (*classService) Update(input classes.Core, id uint) error {
	panic("unimplemented")
}

func New(repo classes.ClassDataInterface) classes.ClassDataInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}
