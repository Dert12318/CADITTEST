package usecaseTest

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"
	"CadItTest/usecase"

	"github.com/gin-gonic/gin"
)

var context *gin.Context

type UsecaseSoal1Struct struct {
	testRepo repository.TestRepoInterface
	log      *log.LogCustom
}

func NewUsecaseImpl(testRepo repository.TestRepoInterface, log *log.LogCustom) usecase.SoalUsecaseInterface {
	return &UsecaseSoal1Struct{testRepo, log}
}

func (b UsecaseSoal1Struct) Soal1Func(input models.DataRequest, id, traceHeader map[string]string) (models.DataRequest, error) {
	url := "http://jsonplaceholder.typicode.com/users"
	if id == nil {
		b.GetDataAllPartner(url, context, traceHeader)
	}
	return input, nil
}
