package delivery

import (
	"immersive-dashboard/features/users"
	helper "immersive-dashboard/utils/helper"
	"log"
	"net/http"

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
