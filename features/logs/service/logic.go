package service

import (
	"immersive-dashboard/features/logs"
	"log"

	"github.com/go-playground/validator/v10"
)

type logService struct {
	data logs.LogData
	vld  *validator.Validate
}

func New(d logs.LogData) logs.LogService {
	return &logService{
		data: d,
		vld:  validator.New(),
	}
}

func (ls *logService) AddLogSrv(newLog logs.Core) error {
	// check input validation
	if errVld := ls.vld.Struct(newLog); errVld != nil {
		log.Println("error validation:", errVld)
		return errVld
	}
	err := ls.data.AddLogData(newLog)
	if err != nil {
		log.Println("error data:", err)
		return err
	}
	return nil
}

func (ls *logService) GetLogSrv(id, pageNum int) ([]logs.Core, error) {
	tmp, err := ls.data.GetLogData(id, pageNum)
	if err != nil {
		log.Println("error data:", err)
		return nil, err
	}
	return tmp, nil
}