package service

import (
	"immersive-dashboard/features/users"
	helper"immersive-dashboard/utils/helper"
	"log"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	data users.UserData
	vld  *validator.Validate
}

func New(d users.UserData, v *validator.Validate) users.UserService {
	return &userService{
		data: d,
		vld:  v,
	}
}

func (us *userService) RegisterSrv(newUser users.Core) error {
	// check input validation
	if errVld := us.vld.Struct(newUser); errVld != nil {
		log.Println("error validation", errVld)
		return errVld
	}
	// bcrypt password before insert into database
	passBcrypt, errBcrypt := helper.PassBcrypt(newUser.Password)
	if errBcrypt != nil {
		log.Println("error bcrypt", errBcrypt)
		return errBcrypt
	}
	
	newUser.Password = passBcrypt

	err := us.data.RegisterData(newUser)
	if err != nil {
		log.Println("error data", err)
		return err
	}

	return nil
}
