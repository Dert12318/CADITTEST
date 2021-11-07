package router

import (
	"CadItTest/config/env"
	"CadItTest/config/log"
	"CadItTest/controller"
	repo1 "CadItTest/repository/RepoSoal"
	usecase1 "CadItTest/usecase/Soal"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	logger := log.NewLogCustom(env.Config)

	// if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
	// 	logger.Fatal(err, "router : init contract", nil)
	// }

	routerGroup := router.Group("/CADIT")
	//repo
	rp := repo1.NewConnectUrlImpl(logger)
	//usecase
	usecaseTest := usecase1.NewUsecaseSoal1Impl(rp, logger)
	//controller
	controller.NewTestController(routerGroup, usecaseTest, logger)
	return router
}
