package delivery

import (
	"errors"
	"immersive-dashboard/features/users"
	jwt "immersive-dashboard/middlewares"
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

// UpdatePass implements users.UserDelivery
func (*userDelivery) UpdatePass() echo.HandlerFunc {
	panic("unimplemented")
}

func New(service users.UserService) users.UserDelivery {
	return &userDelivery{
		srv: service,
	}
}

func (ud *userDelivery) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		if int(jwt.ExtractTokenUserId(c)) != 1 {
			err := errors.New("restricted access")
			return c.JSON(helper.ErrorResponse(err))
		}
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
		return c.JSON(helper.SuccessResponse(http.StatusCreated, "Success Add User"))
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

		res := MentorResponse{}
		copier.Copy(&res, &data)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "Login Success", res, token))
	}
}

func (ud *userDelivery) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		data, err := ud.srv.ProfileSrv(id)
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		res := UserResponse{}
		copier.Copy(&res, &data)
		return c.JSON(helper.SuccessResponse(http.StatusOK, "Success Get Profile", res))
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

		return c.JSON(helper.SuccessResponse(http.StatusOK, "Success Get All User", res))
	}
}

func (ud *userDelivery) GetMentor() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := ud.srv.GetMentorSrv()
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		res := ListMentorResponse{}
		copier.Copy(&res, &data)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "Success Get All Mentor", res))
	}
}

func (ud *userDelivery) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := int(jwt.ExtractTokenUserId(c))
		paramID, _ := strconv.Atoi(c.Param("id"))
		var userID int
		if c.Param("id") != "" && id == 1 {
			userID, _ = strconv.Atoi(c.Param("id"))
		} else if c.Param("id") == "" {
			userID = id
		} else if c.Param("id") != "" && id == paramID {
			userID = paramID
		} else {
			err := errors.New("restricted access")
			return c.JSON(helper.ErrorResponse(err))
		}

		updateInput := UpdateReq{}
		if err := c.Bind(&updateInput); err != nil {
			log.Println("error bind", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		updateUser := users.Core{}
		copier.Copy(&updateUser, &updateInput)
		if err := ud.srv.UpdateUserSrv(userID, updateUser); err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(helper.SuccessResponse(http.StatusOK, "Success Update User"))
	}
}

func (ud *userDelivery) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := strconv.Atoi(c.Param("id"))
		id := int(jwt.ExtractTokenUserId(c))
		if id != 1 {
			err := errors.New("restricted access")
			return c.JSON(helper.ErrorResponse(err))
		}
		if userID == 1 {
			err := errors.New("deleted admin")
			return c.JSON(helper.ErrorResponse(err))
		}
		if err := ud.srv.DeleteSrv(userID); err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(helper.SuccessResponse(http.StatusOK, "Success Delete User"))
	}
}
