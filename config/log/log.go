package log

import (
	"CadItTest/models"
	"fmt"
	"sync"

	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

type LogCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
}

type iAm struct {
	Name string
	Host string
	Port string
}

var instance *LogCustom
var once sync.Once

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom(configServer models.ServerConfig) *LogCustom {
	var log *logrus.Logger

	configElstc := configServer.ElasticConfig

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	client, err := elastic.NewClient(elastic.SetURL(
		fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(configElstc.User, configElstc.Password))
	fmt.Println("http://"+configElstc.Host+":"+configElstc.Port, "<<< ---- elasticsearch url")
	if err != nil {
		selfLogError(err, "config/log: elastic client", log)
	} else {
		hook, err := elogrus.NewAsyncElasticHook(
			client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
		if err != nil {
			selfLogError(err, "config/log: elastic client sync with logrus", log)
		}
		log.Hooks.Add(hook)
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: configServer.Name,
				Host: configServer.Host,
				Port: configServer.Port,
			},
		}
	})
	return instance
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Error(err error, description string, traceHeader map[string]string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

func (l *LogCustom) Success(req interface{}, resp interface{}, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":       l.WhoAmI,
		"trace_header": traceHeader,
		"request":      req,
		"response":     resp,
	}).Info("SUCCESS")
}

func (l *LogCustom) Info(data interface{}, description string, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":       l.WhoAmI,
		"trace_header": traceHeader,
		"data":         data,
	}).Info(description)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Fatal(err error, description string, traceHeader map[string]string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogFatal(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}
