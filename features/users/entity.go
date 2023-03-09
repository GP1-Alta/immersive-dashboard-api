package users

import "github.com/labstack/echo/v4"

type Core struct {
	Id       uint
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string
	Team     string `validate:"required"`
	Status   string `validate:"required"`
}

type UserDelivery interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	GetMentor() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	RegisterSrv(newUser Core) error
	LoginSrv(email, password string) (string, Core, error)
	ProfileSrv(id int)(Core, error)
	GetUser(int, string)([]Core, error)
	GetMentorSrv()([]Core, error)
	UpdateUserSrv(id int, newUser Core) error
	DeleteSrv(id int) error
}

type UserData interface {
	RegisterData(newUser Core) error
	LoginData(email string) (Core, error)
	ProfileData(id int)(Core, error)
	GetUser(int, string)([]Core, error)
	GetMentorData()([]Core, error)
	UpdateUserData(id int, newUser Core) error
	DeleteData(id int) error
}