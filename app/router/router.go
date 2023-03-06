package router

import (
	_menteeData "immersive-dashboard/features/mentees/data"
	_menteeHandler "immersive-dashboard/features/mentees/delivery"
	_menteeService "immersive-dashboard/features/mentees/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	//mentees
	menteeData := _menteeData.New(db)
	menteeService := _menteeService.New(menteeData)
	menteeHandlerAPI := _menteeHandler.New(menteeService)
	e.POST("/mentees", menteeHandlerAPI.Create)
	e.PUT("/mentees/:id", menteeHandlerAPI.Edit)

}
