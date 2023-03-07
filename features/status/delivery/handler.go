package delivery

import (
	"immersive-dashboard/features/status"
	"immersive-dashboard/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	statusService status.StatusServiceInterface
}

func New(srv status.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		statusService: srv,
	}
}

func (delivery *StatusHandler) List(c echo.Context) error {
	data, err := delivery.statusService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := listCoreToResponse(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}
