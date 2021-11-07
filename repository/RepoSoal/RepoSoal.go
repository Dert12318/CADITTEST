package reposoal1

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var context *gin.Context

type ConnectUrlStruct struct {
	log *log.LogCustom
}

func NewConnectUrlImpl(log *log.LogCustom) repository.TestRepoInterface {
	return &ConnectUrlStruct{log}
}

func (b ConnectUrlStruct) GetDataAllPartner(url string, traceHeader map[string]string) ([]models.TestRequest, error) {
	//Get Data from URL
	client := resty.New()
	var data []models.TestRequest
	resp, err := client.R().EnableTrace().SetResult(&data).Get(url)
	if err != nil {
		b.log.Error(err, "Repo: cannot get data from url", traceHeader)
		return []models.TestRequest{}, err
	}
	//Unmarshal Data
	result := resp.String()
	errs := json.Unmarshal([]byte(result), &data)
	if errs != nil {
		b.log.Error(err, "Repo: cannot unmarshal data", traceHeader)
		return []models.TestRequest{}, err
	}
	fmt.Println(data, "lokasi data url")
	return data, nil
}
func (b ConnectUrlStruct) GetDataLokalAllPartner(traceHeader map[string]string) (models.SalaryDataPrtner, error) {
	//Get Data from Lokal
	var file = models.SalaryDataPrtner{}
	data, err := ioutil.ReadFile("salary_data.json")
	if err != nil {
		b.log.Error(err, "Repo: cannot connect data lokal json", traceHeader)
		return models.SalaryDataPrtner{}, err
	}
	//Unmarshal Data
	json.Unmarshal([]byte(data), &file)
	return file, nil
}

func (b ConnectUrlStruct) GetDataSensor(traceHeader map[string]string) ([]models.DataSensor, error) {
	//Get Data from Lokal
	var file = []models.DataSensor{}
	data, err := ioutil.ReadFile("sensor_data.json")
	if err != nil {
		b.log.Error(err, "Repo: cannot connect data lokal json", traceHeader)
		return []models.DataSensor{}, err
	}
	//Unmarshal Data
	json.Unmarshal([]byte(data), &file)
	return file, nil
}
