package usecase

import (
	"CadItTest/models"
)

type SoalUsecaseInterface interface {
	Soal1Func(id string, traceHeader map[string]string) (models.DataRequest, error)
}
