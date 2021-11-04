package env

import (
	"CadItTest/models"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var Config models.ServerConfig

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func init() {
	err := loadConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error_cause":   PrintErrorStack(err),
			"error_message": err.Error(),
		}).Fatal("config/env: load config")
	}
}

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal(err, " config/env: load gotdotenv")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.ElasticConfig)
	if err != nil {
		return err
	}

	return err
}

func PrintErrorStack(err error) string {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	return stFormat
}
