package usecase

import (
	"CadItTest/models"
)

type SoalUsecaseInterface interface {
	Soal1Func1(traceHeader map[string]string) ([]models.DataResponse, error)
	Soal1Func2(id string, traceHeader map[string]string) (models.DataResponse, error)
	Soal2Func2(byDay string, roomArea string, traceHeader map[string]string) (models.DataSensorRespone, error)
}
