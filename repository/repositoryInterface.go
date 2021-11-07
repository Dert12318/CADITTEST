package repository

import (
	"CadItTest/models"
)

type TestRepoInterface interface {
	GetDataAllPartner(url string, traceHeader map[string]string) ([]models.TestRequest, error)
	GetDataLokalAllPartner(traceHeader map[string]string) (models.SalaryDataPrtner, error)
	GetDataSensor(traceHeader map[string]string) ([]models.DataSensor, error)
}
