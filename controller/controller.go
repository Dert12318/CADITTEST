package controller

import (
	"CadItTest/config/log"
	"CadItTest/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	SoalUsecase usecase.SoalUsecaseInterface
	log         *log.LogCustom
}

func NewTestController(r *gin.RouterGroup, TestUsecase usecase.SoalUsecaseInterface, log *log.LogCustom) {
	handler := TestController{TestUsecase, log}

	r.GET("/AllPartner", handler.Soal1Test)
}

func (c TestController) Soal1Test(context *gin.Context) {
	var id = context.Query("id")
	fmt.Println(id)
	response, err := c.SoalUsecase.Soal1Func(id, nil)
	if err != nil {
		context.JSON(http.StatusBadRequest, "data error")
		return
	}
	context.JSON(http.StatusOK, response)

}
