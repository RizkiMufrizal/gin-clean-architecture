package common

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		exception.PanicLogging(err)
	}

	date := time.Now()
	logFile, err := os.OpenFile("logs/log_"+date.Format("01-02-2006_15")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicLogging(err)
	if err == nil {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		logger.SetOutput(multiWriter)
	}
	return logger
}

type LogrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter(logger *logrus.Logger) *LogrusAdapter {
	return &LogrusAdapter{logger: logger}
}

func (l *LogrusAdapter) Log(ctx context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {
	entry := l.logger.WithContext(ctx).WithFields(data)

	switch level {
	case sqldblogger.LevelError:
		entry.Error(msg)
	case sqldblogger.LevelInfo:
		entry.Info(msg)
	case sqldblogger.LevelDebug:
		entry.Debug(msg)
	case sqldblogger.LevelTrace:
		entry.Trace(msg)
	default:
		entry.Debug(msg)
	}
}
