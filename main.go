package main

import (
	"CadItTest/config/env"
	"CadItTest/router"
)

func main() {
	StartApp()
}

func StartApp() {
	addr := env.Config.Host + ":" + env.Config.Port
	router.NewRouter().Run(addr)
}
