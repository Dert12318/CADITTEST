package Soal1

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

func (b UsecaseSoal1Struct) Soal1Func1(id string, traceHeader map[string]string) ([]models.TestRequest, error) {
	url := "http://jsonplaceholder.typicode.com/users"
	res, err := b.testRepo.GetDataAllPartner(url, traceHeader)
	if err != nil {
		b.log.Error(err, "usecase/Soal1: cnnot execute get data from repo", traceHeader)
		return []models.TestRequest{}, err
	}
	return res, nil
}
func (b UsecaseSoal1Struct) Soal1Func2(id string, traceHeader map[string]string) ([]models.TestRequest, error) {
	url := "http://jsonplacehoder.typicode.com/users"
	urlku := url + "?id=" + id
	res, err := b.testRepo.GetDataAllPartner(urlku, traceHeader)
	if err != nil {
		b.log.Error(err, "usecase/Soal1: cnnot execute get data from repo", traceHeader)
		return []models.TestRequest{}, err
	}
	return res, nil
}

//
