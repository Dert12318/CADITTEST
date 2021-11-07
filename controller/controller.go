package controller

import (
	"CadItTest/config/log"
	"CadItTest/usecase"
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
	r.GET("/AgregateData", handler.Soal1Test)
}

func (c TestController) Soal1Test(context *gin.Context) {
	var id = context.Query("id")
	if id == "" {
		response, err := c.SoalUsecase.Soal1Func1(nil)
		if err != nil {
			context.JSON(http.StatusBadRequest, "data error")
			return
		}
		context.JSON(http.StatusOK, response)
	} else {
		response2, err2 := c.SoalUsecase.Soal1Func2(id, nil)
		if err2 != nil {
			context.JSON(http.StatusBadRequest, "data error")
			return
		}
		context.JSON(http.StatusOK, response2)
	}
}
func (c TestController) Soal2Test(context *gin.Context) {
	var byDay = context.Query("byDay")
	var roomArea = context.Query("roomArea")
	response, err2 := c.SoalUsecase.Soal2Func2(byDay, roomArea, nil)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, "data error")
		return
	}
	context.JSON(http.StatusOK, response)
}
