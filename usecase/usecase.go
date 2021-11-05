package usecase

import (
	"CadItTest/models"
)

type SoalUsecaseInterface interface {
	Soal1Func1(id string, traceHeader map[string]string) ([]models.TestRequest, error)
	Soal1Func2(id string, traceHeader map[string]string) ([]models.TestRequest, error)
}
