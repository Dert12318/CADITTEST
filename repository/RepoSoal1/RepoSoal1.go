package reposoal1

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"
	"encoding/json"
	"fmt"

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
	fmt.Println(url)
	resp, err := client.R().EnableTrace().SetResult(&data).Get(url)
	if err != nil {
		b.log.Error(err, "Repo: cannot get data from url", traceHeader)
		return []models.TestRequest{}, err
	}
	//Unmarshal Data
	result := resp.String()
	fmt.Println(err, resp)
	errs := json.Unmarshal([]byte(result), &data)
	if errs != nil {
		b.log.Error(err, "Repo: cannot unmarshal data", traceHeader)
		return []models.TestRequest{}, err
	}
	return data, nil
}
