package usecaseTest

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"
	"CadItTest/usecase"
)
var context *gin.Context
type UsecaseSoal1Struct struct {
	testRepo repository.TestRepoInterface
	log      *log.LogCustom
}

func NewUsecaseImpl(TestRepo repository.TestRepoInterface, log *log.LogCustom) usecase.SoalUsecaseInterface {
	return &UsecaseSoal1Struct{TestRepo, log}
}

func (b UsecaseSoal1Struct) Soal1Func(input models.DataRequest, id, traceHeader map[string]string) (models.DataRequest, error) {
	url := "http://jsonplaceholder.typicode.com/users"
	if (id==nil){
		b.GetDataAllPartner(url, context, traceHeader,)
	}
	return input, nil
}
