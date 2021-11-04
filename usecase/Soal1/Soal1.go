package usecaseTest

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"
	"CadItTest/usecase"
)

type UsecaseSoal1Struct struct {
	testRepo repository.TestRepoInterface
	log      *log.LogCustom
}

func NewUsecaseSoal1Impl(testRepo repository.TestRepoInterface, log *log.LogCustom) usecase.SoalUsecaseInterface {
	return &UsecaseSoal1Struct{testRepo, log}
}

func (b UsecaseSoal1Struct) Soal1Func(id string, traceHeader map[string]string) (models.DataRequest, error) {
	url := "http://jsonplaceholder.typicode.com/users"
	res, err := b.testRepo.GetDataAllPartner(url, traceHeader)
	if err != nil {
		b.log.Error(err, "usecase/Soal1: cannot execute get data from repo", traceHeader)
		return models.DataRequest{}, err
	}
	return res, nil
}
