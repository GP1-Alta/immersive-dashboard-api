package logs

import (
	"github.com/labstack/echo/v4"
	// "gorm.io/datatypes"
)

type Core struct {
	ID         uint
	MenteeID   uint
	UserID     uint
	UserName   string
	StatusID   uint
	StatusName string
	Feedback   string `validate:"required"`
	CreatedAt  string
	Image      string
}

type LogDelivery interface {
	AddLog() echo.HandlerFunc
	GetLog() echo.HandlerFunc
}

type LogService interface {
	AddLogSrv(newLog Core) error
	GetLogSrv(id, page int) ([]Core, error)
}

type LogData interface {
	AddLogData(newLog Core) error
	GetLogData(id, page int) ([]Core, error)
}
