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
	"immersive-dashboard/middlewares"

	_logData "immersive-dashboard/features/logs/data"
	_logHandler "immersive-dashboard/features/logs/delivery"
	_logService "immersive-dashboard/features/logs/service"
	_userData "immersive-dashboard/features/users/data"
	_userHandler "immersive-dashboard/features/users/delivery"
	_userService "immersive-dashboard/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	//mentees
	menteeData := _menteeData.New(db)
	menteeService := _menteeService.New(menteeData)
	menteeHandlerAPI := _menteeHandler.New(menteeService)
	e.GET("/mentees", menteeHandlerAPI.GetAll, middlewares.JWTMiddleware())
	e.GET("/mentees/:id", menteeHandlerAPI.Get, middlewares.JWTMiddleware())
	e.POST("/mentees", menteeHandlerAPI.Create, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", menteeHandlerAPI.Edit, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", menteeHandlerAPI.Delete, middlewares.JWTMiddleware())

	//classes
	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)
	e.POST("/classes", classHandlerAPI.Create, middlewares.JWTMiddleware())
	e.GET("/classes", classHandlerAPI.GetAll, middlewares.JWTMiddleware())
	e.GET("/classes/list", classHandlerAPI.List, middlewares.JWTMiddleware())
	e.PUT("/classes/:id", classHandlerAPI.Edit, middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", classHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.GET("/classes/:id", classHandlerAPI.Detail, middlewares.JWTMiddleware())

	//status
	statusData := _statusData.New(db)
	statusService := _statusService.New(statusData)
	statusHandlerAPI := _statusHandler.New(statusService)
	e.GET("/status", statusHandlerAPI.List, middlewares.JWTMiddleware())

	userData := _userData.New(db)
	userSrv := _userService.New(userData)
	userHdl := _userHandler.New(userSrv)
	e.POST("/login", userHdl.Login())
	e.POST("/register", userHdl.Register(), middlewares.JWTMiddleware())
	e.GET("/users", userHdl.GetUser(), middlewares.JWTMiddleware())
	e.GET("/mentors", userHdl.GetMentor(), middlewares.JWTMiddleware())
	e.GET("/profile/:id", userHdl.Profile(), middlewares.JWTMiddleware())
	e.PUT("/users", userHdl.UpdateUser(), middlewares.JWTMiddleware())
	e.PUT("/users/:id", userHdl.UpdateUser(), middlewares.JWTMiddleware())
	e.DELETE("/users/:id", userHdl.Delete(), middlewares.JWTMiddleware())

	logData := _logData.New(db)
	logSrv := _logService.New(logData)
	logHdl := _logHandler.New(logSrv)
	e.POST("/mentees/:id/logs", logHdl.AddLog(), middlewares.JWTMiddleware())
	e.GET("/mentees/:id/logs", logHdl.GetLog(), middlewares.JWTMiddleware())
}
