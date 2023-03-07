package data

import (
	"immersive-dashboard/features/logs"
	"log"

	"gorm.io/gorm"
)

type logQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) logs.LogData {
	return &logQuery{
		db: db,
	}
}

func (lq *logQuery) AddLogData(newLog logs.Core) error {
	data := CoreToLog(newLog)
	tx := lq.db.Create(&data)
	if tx.Error != nil {
		log.Println("error query", tx.Error)
		return tx.Error
	}
	return nil
}