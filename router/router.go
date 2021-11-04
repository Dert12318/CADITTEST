package router

import (
	// "CadItTest/config/env"
	// "CadItTest/config/log"
	// "CadItTest/controller"
	// "CadItTest/repository"
	//usecaseTest "CadItTest/usecase/usecaseTest"

	//"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	// logger := log.NewLogCustom(env.Config)

	// if err := errcntrct.InitContract(env.Config.JSONPathFile); err != nil {
	// 	logger.Fatal(err, "router : init contract", nil)
	// }

	// routerGroup := router.Group("CADIT/")

	// //repo
	// repo := connectesb.NewConnectESBRepoImpl()
	// //usecase
	// usecaseTest := usecaseTest.NewTestStructImpl(repo , logger)
	// //controller
	// controller.NewTestController(routerGroup, usecaseTest, logger)
	return router
}
