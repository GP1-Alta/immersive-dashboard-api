package router

import (
	_classData "immersive-dashboard/features/classes/data"
	_classHandler "immersive-dashboard/features/classes/delivery"
	_classService "immersive-dashboard/features/classes/service"
	_menteeData "immersive-dashboard/features/mentees/data"
	_menteeHandler "immersive-dashboard/features/mentees/delivery"
	_menteeService "immersive-dashboard/features/mentees/service"
	_statusData "immersive-dashboard/features/status/data"
	_statusHandler "immersive-dashboard/features/status/delivery"
	_statusService "immersive-dashboard/features/status/service"

	_logData "immersive-dashboard/features/logs/data"
	_logHandler "immersive-dashboard/features/logs/delivery"
	_logService "immersive-dashboard/features/logs/service"
	_userData "immersive-dashboard/features/users/data"
	_userHandler "immersive-dashboard/features/users/delivery"
	_userService "immersive-dashboard/features/users/service"

	"github.com/go-playground/validator/v10"
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

	//classes
	statusData := _statusData.New(db)
	statusService := _statusService.New(statusData)
	statusHandlerAPI := _statusHandler.New(statusService)
	e.GET("/status", statusHandlerAPI.List)

	v := validator.New()
	userData := _userData.New(db)
	userSrv := _userService.New(userData, v)
	userHdl := _userHandler.New(userSrv)

	logData := _logData.New(db)
	logSrv := _logService.New(logData, v)
	logHdl := _logHandler.New(logSrv)

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.GetUser())
	e.PUT("/users/:id", userHdl.UpdateUser())
	e.DELETE("/users/:id", userHdl.Delete())

	e.POST("/mentees/:mentee_id/logs", logHdl.AddLog())
}
