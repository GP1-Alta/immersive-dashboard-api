package delivery

import (
	"immersive-dashboard/features/users"
	helper "immersive-dashboard/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	srv users.UserService
}

func New(service users.UserService) users.UserDelivery {
	return &userDelivery{
		srv: service,
	}
}

func (ud *userDelivery) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		regInput := RegisterReq{}
		if err := c.Bind(&regInput); err != nil {
			log.Println("error bind", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		newUser := users.Core{}
		copier.Copy(&newUser, &regInput)
		err := ud.srv.RegisterSrv(newUser)
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(helper.SuccessResponse(http.StatusCreated, "Success register account"))
	}
}

func (ud *userDelivery) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginInput := LoginReq{}
		if err := c.Bind(&loginInput); err != nil {
			log.Println("error bind", err)
			return c.JSON(helper.ErrorResponse(err))
		}

		token, data, err := ud.srv.LoginSrv(loginInput.Email, loginInput.Password)
		if err != nil {
			return c.JSON(helper.ErrorResponse(err))
		}

		res := UserResponse{}
		copier.Copy(&res, &data)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "Login success", res, token))
	}
}

func (ud *userDelivery) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		key := c.QueryParam("key")
		data, err := ud.srv.GetUser(page, key)
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		res := ListUserResponse{}
		copier.Copy(&res, &data)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "successfully get all user", res))
	}
}
