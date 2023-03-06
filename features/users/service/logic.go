package service

import (
	"errors"
	"immersive-dashboard/features/users"
	"immersive-dashboard/middlewares"
	helper "immersive-dashboard/utils/helper"
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
		log.Println("error validation:", errVld)
		return errVld
	}
	// bcrypt password before insert into database
	passBcrypt, errBcrypt := helper.PassBcrypt(newUser.Password)
	if errBcrypt != nil {
		log.Println("error bcrypt:", errBcrypt)
		return errBcrypt
	}

	newUser.Password = passBcrypt

	err := us.data.RegisterData(newUser)
	if err != nil {
		log.Println("error data:", err)
		return err
	}

	return nil
}

// LoginSrv implements users.UserService
func (us *userService) LoginSrv(email string, password string) (string, users.Core, error) {
	if email == "" {
		return "", users.Core{}, errors.New("email cannot be empty")
	} else if password == "" {
		return "", users.Core{}, errors.New("password cannot be empty")
	}

	tmp, err := us.data.LoginData(email)
	if err != nil {
		log.Println("error service:", err)
		return "", users.Core{}, err
	}

	if err := helper.PassCompare(tmp.Password, password); err != nil {
		log.Println("error compare password:", err)
		return "", users.Core{}, err
	}

	token, err := middlewares.CreateToken(tmp.Id)
	if err != nil {
		log.Println("error generate token:", err)
		return "", users.Core{}, err
	}

	return token, tmp, nil
}