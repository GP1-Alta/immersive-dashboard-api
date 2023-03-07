package data

import (
	"immersive-dashboard/features/status"

	"gorm.io/gorm"
)

type statusQuery struct {
	db *gorm.DB
}

// SelectAll implements status.StatusDataInterface
func (repo *statusQuery) SelectAll() ([]status.Core, error) {
	var statusModel []Status
	tx := repo.db.Find(&statusModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	statusCoreAll := ListModelToCore(statusModel)
	return statusCoreAll, nil
}

func New(db *gorm.DB) status.StatusDataInterface {
	return &statusQuery{
		db: db,
	}
}
