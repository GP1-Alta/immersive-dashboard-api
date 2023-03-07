package router

import (
	_classData "immersive-dashboard/features/classes/data"
	_classHandler "immersive-dashboard/features/classes/delivery"
	_classService "immersive-dashboard/features/classes/service"
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
	e.GET("/mentees", menteeHandlerAPI.GetAll)
	e.POST("/mentees", menteeHandlerAPI.Create)
	e.PUT("/mentees/:id", menteeHandlerAPI.Edit)
	e.DELETE("/mentees/:id", menteeHandlerAPI.Delete)

	//classes
	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)
	e.POST("/classes", classHandlerAPI.Create)
	e.GET("/classes", classHandlerAPI.GetAll)
	e.GET("/classes/list", classHandlerAPI.List)
	e.PUT("/classes/:id", classHandlerAPI.Edit)
	e.DELETE("/classes/:id", classHandlerAPI.Delete)
}
