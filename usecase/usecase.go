package usecase

import (
	"CadItTest/models"
)

type SoalUsecaseInterface interface {
	Soal1Func(input models.DataRequest, id string, traceHeader map[string]string) (models.DataRequest, error)
}
