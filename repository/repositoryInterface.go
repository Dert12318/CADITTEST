package repository

import (
	"CadItTest/models"
)

type TestRepoInterface interface {
	GetDataAllPartner(url string, traceHeader map[string]string) ([]models.TestRequest, error)
}
