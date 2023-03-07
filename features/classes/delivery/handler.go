package delivery

import (
	"immersive-dashboard/features/classes"
	"immersive-dashboard/utils/helper"
	"net/http"
	"strconv"

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

func (delivery *ClassHandler) GetAll(c echo.Context) error {
	var pageNumber int = 1
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam)
		if errConv != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response("Failed, page must number"))
		} else {
			pageNumber = pageConv
		}
	}

	nameParam := c.QueryParam("name")
	data, err := delivery.classService.GetAll(pageNumber, nameParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := listCoreToResponse(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}
