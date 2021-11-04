package repository

import (
	"github.com/gin-gonic/gin"
	"CadItTest/models"
)

type TestRepoInterface interface {
	GetDataAllPartner(url string, context *gin.Context, traceHeader map[string]string) (models.DataRequest, error)
}
