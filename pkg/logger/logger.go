package logger

import (
	"strings"

	"github.com/challenge/pkg/config"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Setup() {
	logLevel := strings.ToUpper(config.GetConfig().LogLevel)

	switch logLevel {
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		log.SetLevel(logrus.InfoLevel)
	case "WARNING":
		log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}

func Trace(msg string) {
	log.Trace(msg)
}

func Tracef(msg string, args ...interface{}) {
	log.Tracef(msg, args...)
}

func TraceWithFields(msg string, fields map[string]interface{}) {
	log.WithFields(logrus.Fields(fields)).Trace(msg)
}

func Debug(msg string) {
	log.Debug(msg)
}

func Debugf(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func DebugWithFields(msg string, fields map[string]interface{}) {
	log.WithFields(logrus.Fields(fields)).Debug(msg)
}

func Info(msg string) {
	log.Info(msg)
}

func Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func InfoWithFields(msg string, fields map[string]interface{}) {
	log.WithFields(logrus.Fields(fields)).Info(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Warnf(msg string, args ...interface{}) {
	log.Warnf(msg, args...)
}

func WarnWithFields(msg string, fields map[string]interface{}) {
	log.WithFields(logrus.Fields(fields)).Warn(msg)
}

func Error(msg string) {
	log.Error(msg)
}

func Errorf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

func ErrorWithFields(msg string, fields map[string]interface{}) {
	log.WithFields(logrus.Fields(fields)).Error(msg)
}

func Fatal(msg string) {
	log.Fatal(msg)
}

func Fatalf(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}

func FatalWithFields(msg string, fields map[string]interface{}) {
	log.WithFields(logrus.Fields(fields)).Fatal(msg)
}
