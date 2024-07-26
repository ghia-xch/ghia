package util

import "github.com/sirupsen/logrus"

func GetLogLevel(level string, defaultLevel logrus.Level) logrus.Level {

	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return defaultLevel
	}

}
