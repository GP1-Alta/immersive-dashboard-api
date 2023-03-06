package delivery

import (
	"immersive-dashboard/features/mentees"
	"immersive-dashboard/utils/helper"
	"net/http"

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

func (delivery *MenteeHandler) CreateMentee(c echo.Context) error {
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
