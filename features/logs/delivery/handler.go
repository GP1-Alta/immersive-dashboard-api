package delivery

import (
	"immersive-dashboard/features/logs"
	jwt "immersive-dashboard/middlewares"
	helper "immersive-dashboard/utils/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type logDelivery struct {
	srv logs.LogService
}

func New(service logs.LogService) logs.LogDelivery {
	return &logDelivery{
		srv: service,
	}
}

func (ld *logDelivery) AddLog() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := int(jwt.ExtractTokenUserId(c))
		menteeID, _ := strconv.Atoi(c.Param("id"))

		addInput := AddLogReq{}
		if err := c.Bind(&addInput); err != nil {
			log.Println("error bind", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		addInput.MenteeID = uint(menteeID)
		addInput.UserID = uint(userID)

		newLog := logs.Core{}
		copier.Copy(&newLog, &addInput)

		err := ld.srv.AddLogSrv(newLog)
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(helper.SuccessResponse(http.StatusCreated, "Success Add Log"))
	}
}

func (ld *logDelivery) GetLog() echo.HandlerFunc {
	return func(c echo.Context) error {
		menteeID, _ := strconv.Atoi(c.Param("id"))
		page, _ := strconv.Atoi(c.QueryParam("page"))
		data, err := ld.srv.GetLogSrv(menteeID, page)
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		res := ListLogResponse{}
		copier.Copy(&res, &data)

		return c.JSON(helper.SuccessResponse(http.StatusOK, "Success Get All Log", res))
	}
}
