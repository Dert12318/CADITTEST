package reposoal1

import (
	"CadItTest/config/log"
	"CadItTest/models"
	"CadItTest/repository"

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

func (b ConnectUrlStruct) GetDataAllPartner(url string, traceHeader map[string]string) (models.DataRequest, error) {
	//Get Data from URL
	client := resty.New()
	var data models.DataRequest
	_, err := client.R().EnableTrace().SetResult(&data).Get(url)
	if err != nil {
		b.log.Error(err, "Repo: cannot get data from url", traceHeader)
		return models.DataRequest{}, err
	}
	return data, nil
}

// func (b ConnectUrlStruct) GetDataPartnerId(input models.AccessData, context *gin.Context, param id, traceHeader map[string]string) (models.DataRequest, error) {
// 	//Get Data from URL
// 	client := resty.New()
// 	var data models.DataRequest
// 	_, err := client.R().EnableTrace().SetResult(&data).Get(input.UrlSoal1)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, "data error binding")
// 		return models.DataRequest{},err
// 	}
// 	return data, nil
// }
