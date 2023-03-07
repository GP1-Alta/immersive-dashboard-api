package logs

import "github.com/labstack/echo/v4"

type Core struct {
	Id       uint
	MenteeID uint
	UserID   uint
	StatusID uint
	Feedback string `validate:"required"`
}

type LogDelivery interface {
	AddLog() echo.HandlerFunc
}

type LogService interface {
	AddLogSrv(newLog Core) error
}

type LogData interface {
	AddLogData(newLog Core) error
}