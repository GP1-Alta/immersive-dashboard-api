package delivery

import (
	"immersive-dashboard/features/mentees"
	"immersive-dashboard/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	menteeService mentees.MenteeServiceInterface
}

func New(srv mentees.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: srv,
	}
}

func (delivery *MenteeHandler) Create(c echo.Context) error {
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Response("Failed Add Mentee, error bind data"))
	}
	dataCore := requestToCore(menteeInput)
	err := delivery.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed Add Mentee, error: "+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.Response("Success Add Mentee"))
}

func (delivery *MenteeHandler) Edit(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.Response("Failed Update Mentee, id param must number"))
	}
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.Response("Failed Update Mentee, error bind data"))
	}

	dataCore := requestToCore(menteeInput)
	err := delivery.menteeService.Edit(dataCore, uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed Update Mentee, error update data"))
	}
	return c.JSON(http.StatusOK, helper.Response("Success Update Mentee"))
}

func (delivery *MenteeHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.Response("Failed Delete Mentee, id param must number"))
	}
	var dataCore mentees.Core
	err := delivery.menteeService.Delete(dataCore, uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed Delete Mentee, error delete data"))
	}
	return c.JSON(http.StatusOK, helper.Response("Success Delete Mentee"))
}

func (delivery *MenteeHandler) GetAll(c echo.Context) error {
	data, err := delivery.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := listCoreToResponse(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}
