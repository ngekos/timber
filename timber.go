package timber

import (
	"os"

	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type (
	//Woodman ...
	Woodman struct {
		name   string
		logger *log.Logger
	}

	//Timber ...
	Timber interface {
		Logger() *log.Logger
		Name() string

		LogDebug(...interface{})
		LogInfo(...interface{})
		LogError(...interface{})
	}
)

//NewTimber ...
func NewTimber(name string, logger *log.Logger) *Woodman {
	return &Woodman{
		name:   name,
		logger: logger,
	}
}

//Logger ...
func (o *Woodman) Logger() *log.Logger {
	return o.logger
}

//Name ...
func (o *Woodman) Name() string {
	return o.name
}

//LogDebug ...
func (o *Woodman) LogDebug(args ...interface{}) {
	o.logger.Debugln(args)
}

//LogInfo ...
func (o *Woodman) LogInfo(args ...interface{}) {
	o.logger.Infoln(args)
}

//LogError ...
func (o *Woodman) LogError(args ...interface{}) {
	o.logger.Errorln(args)
}

//NewDefaultLogger because log is dependency
func NewDefaultLogger() *log.Logger {
	logger := log.New()
	logger.Out = os.Stdout
	logger.Formatter = &log.TextFormatter{}
	return logger
}

//NewLumberjackLogger because log is dependency
func NewLumberjackLogger(filename string, maxsize, maxbackup, maxage int) *log.Logger {
	logger := log.New()
	logger.Out = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,   //megabytes
		MaxBackups: maxbackup, //number
		MaxAge:     maxage,    //days
	}
	logger.Formatter = &log.JSONFormatter{}
	return logger
}

//NewJSONLogger for JSON stdout
func NewJSONLogger() *log.Logger {
	logger := NewDefaultLogger()
	logger.Formatter = &log.JSONFormatter{}
	return logger
}
