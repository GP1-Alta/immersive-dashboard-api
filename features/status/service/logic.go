package service

import (
	"immersive-dashboard/features/status"

	"github.com/go-playground/validator/v10"
)

type statusService struct {
	statusData status.StatusDataInterface
	validate   *validator.Validate
}

// GetAll implements status.StatusServiceInterface
func (service *statusService) GetAll() ([]status.Core, error) {
	data, err := service.statusData.SelectAll()
	return data, err
}

func New(repo status.StatusDataInterface) status.StatusServiceInterface {
	return &statusService{
		statusData: repo,
		validate:   validator.New(),
	}
}
