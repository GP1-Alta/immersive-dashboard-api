package data

import (
	"errors"
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

func (lq *logQuery) GetLogData(menteeID, pageNum int) ([]logs.Core, error) {
	tmp := []Log{}
	limit := 5
	offset := (pageNum - 1) * limit
	tx := lq.db.Limit(limit).Offset(offset).Where("logs.mentee_id = ?", menteeID).Select("logs.id, logs.created_at, logs.image, logs.feedback, users.name AS user_name, statuses.name AS status_name").Joins("JOIN users ON logs.user_id = users.id").Joins("JOIN statuses ON logs.status_id = statuses.id").Find(&tmp)
	if tx.RowsAffected < 1 {
		return nil, errors.New("not found")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	listUser := ListLogToCore(tmp)
	return listUser, nil
}
