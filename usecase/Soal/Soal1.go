package Soal1

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"
	"CadItTest/usecase"
	"fmt"
	"strconv"
)

type UsecaseSoal1Struct struct {
	testRepo repository.TestRepoInterface
	log      *log.LogCustom
}

func NewUsecaseSoal1Impl(testRepo repository.TestRepoInterface, log *log.LogCustom) usecase.SoalUsecaseInterface {
	return &UsecaseSoal1Struct{testRepo, log}
}

func (b UsecaseSoal1Struct) Soal1Func1(traceHeader map[string]string) ([]models.DataResponse, error) {
	var response = []models.DataResponse{}

	//Get data url
	url := "https://jsonplaceholder.typicode.com/users"
	res, err := b.testRepo.GetDataAllPartner(url, traceHeader)
	if err != nil {
		b.log.Error(err, "usecase/Soal1: cnnot execute get data from repo", traceHeader)
		return []models.DataResponse{}, err
	}
	fmt.Println(res)
	//Get data lokal Salary
	datalokal, err2 := b.testRepo.GetDataLokalAllPartner(traceHeader)
	fmt.Println(datalokal)
	if err2 != nil {
		b.log.Error(err, "usecase/Soal1: cnnot execute get data lokal from repo", traceHeader)
		return []models.DataResponse{}, err
	}
	length := len(datalokal.Array)
	//Parsing Data
	for i := 0; i < length; i++ {
		response[i].ID = res[i].ID
		response[i].NamePartner = res[i].NamePartner
		response[i].UsernamePartner = res[i].UsernamePartner
		response[i].EmailPartner = res[i].EmailPartner
		response[i].AddressPartner.CityPartner = res[i].AddressPartner.CityPartner
		response[i].AddressPartner.GeoPartner = res[i].AddressPartner.GeoPartner
		response[i].AddressPartner.StreetPartner = res[i].AddressPartner.StreetPartner
		response[i].AddressPartner.SuitePartner = res[i].AddressPartner.SuitePartner
		response[i].AddressPartner.ZIPCodePartner = res[i].AddressPartner.ZIPCodePartner
		response[i].AddressPartner.GeoPartner.LatGeoPartner = res[i].AddressPartner.GeoPartner.LatGeoPartner
		response[i].AddressPartner.GeoPartner.LngGeoPartner = res[i].AddressPartner.GeoPartner.LngGeoPartner
		response[i].NamePartner = res[i].NamePartner
		response[i].PhonePartner = res[i].PhonePartner
		response[i].SalaryIDR = datalokal.Array[i].SalaryPartner
		response[i].SalaryUSD = datalokal.Array[i].SalaryPartner
	}
	//Get datalokal kurs IDR to USD
	return response, nil
}
func (b UsecaseSoal1Struct) Soal1Func2(id string, traceHeader map[string]string) (models.DataResponse, error) {
	//Get data url
	response := models.DataResponse{}
	url := "http://jsonplaceholder.typicode.com/users"
	urlku := url + "?id=" + id
	res, err := b.testRepo.GetDataAllPartner(urlku, traceHeader)
	fmt.Println(err, "data url")
	if err != nil {
		b.log.Error(err, "usecase/Soal1: cnnot execute get data from repo", traceHeader)
		return models.DataResponse{}, err
	}
	//Get data lokal Salary
	datalokal, err2 := b.testRepo.GetDataLokalAllPartner(traceHeader)
	fmt.Println(datalokal, "data lokal")
	if err2 != nil {
		b.log.Error(err2, "usecase/Soal1: cnnot execute get data lokal from repo", traceHeader)
		return models.DataResponse{}, err2
	}
	NewId, _ := strconv.Atoi(id)
	//parsing data
	response.ID = res[0].ID
	response.NamePartner = res[0].NamePartner
	response.UsernamePartner = res[0].UsernamePartner
	response.EmailPartner = res[0].EmailPartner
	response.AddressPartner.CityPartner = res[0].AddressPartner.CityPartner
	response.AddressPartner.GeoPartner = res[0].AddressPartner.GeoPartner
	response.AddressPartner.StreetPartner = res[0].AddressPartner.StreetPartner
	response.AddressPartner.SuitePartner = res[0].AddressPartner.SuitePartner
	response.AddressPartner.ZIPCodePartner = res[0].AddressPartner.ZIPCodePartner
	response.AddressPartner.GeoPartner.LatGeoPartner = res[0].AddressPartner.GeoPartner.LatGeoPartner
	response.AddressPartner.GeoPartner.LngGeoPartner = res[0].AddressPartner.GeoPartner.LngGeoPartner
	response.PhonePartner = res[0].PhonePartner
	response.SalaryIDR = datalokal.Array[NewId-1].SalaryPartner
	response.SalaryUSD = datalokal.Array[NewId-1].SalaryPartner
	fmt.Println(response)
	return response, nil
}

//
