package data

import (
	"immersive-dashboard/features/users"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) RegisterData(newUser users.Core) error {
	data := CoreToUser(newUser)
	tx := uq.db.Create(&data)
	if tx.Error != nil {
		log.Println("error query", tx.Error)
		return tx.Error
	}
	return nil
}

func (uq *userQuery) LoginData(email string) (users.Core, error) {
	tmp := User{}
	tx := uq.db.Where("email = ?", email).First(&tmp)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return UserToCore(tmp), nil
}