package delivery

import (
	"immersive-dashboard/features/logs"
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
		menteeID, errCnv := strconv.Atoi(c.Param("mentee_id"))
		if errCnv != nil {
			return c.JSON(helper.ErrorResponse(errCnv))
		}
		
		addInput := AddLogReq{}
		if err := c.Bind(&addInput); err != nil {
			log.Println("error bind", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		addInput.MenteeID = uint(menteeID)
		
		newLog := logs.Core{}
		copier.Copy(&newLog, &addInput)

		err := ld.srv.AddLogSrv(newLog)
		if err != nil {
			log.Println("error handler", err)
			return c.JSON(helper.ErrorResponse(err))
		}
		return c.JSON(helper.SuccessResponse(http.StatusCreated, "Success Added Log"))
	}
}
