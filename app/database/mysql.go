package database

import (
	"fmt"
	"immersive-dashboard/app/config"
	_modelClass "immersive-dashboard/features/classes/data"
	_modelLog "immersive-dashboard/features/logs/data"
	_modelMentee "immersive-dashboard/features/mentees/data"
	_modelStatus "immersive-dashboard/features/status/data"
	_modelUser "immersive-dashboard/features/users/data"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMySql(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to DB:", err.Error())
		return nil
	}

	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_modelUser.User{})
	db.AutoMigrate(&_modelClass.Class{})
	db.AutoMigrate(&_modelStatus.Status{})
	db.AutoMigrate(&_modelMentee.Mentee{})
	db.AutoMigrate(&_modelLog.Log{})
}
