package delivery

import (
	"immersive-dashboard/features/classes"
	"immersive-dashboard/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService classes.ClassServiceInterface
}

func New(srv classes.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: srv,
	}
}

func (delivery *ClassHandler) Create(c echo.Context) error {
	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Response("Failed Add Class, error bind data"))
	}
	dataCore := requestToCore(classInput)
	err := delivery.classService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed Add Class, error: "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.Response("Success Add Class"))
}
