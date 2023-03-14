package global

import (
	"bytes"
	"fmt"
	"inception/api/config"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

type LogrusConfig struct {
	LevelStr     string
	Level        logrus.Level
	LogInConsole bool
	ShowLine     bool
	Filename     string
	MaxSize      int
	MaxBackups   int
	MaxAge       int
	Compress     bool
}

func InitLogrus(config *LogrusConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(config.Level)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logFormatter{})
	logrus.SetLevel(config.Level)
	logger.SetOutput(os.Stdout)
	// if config.Filename != "" {
	// 	logrus.SetOutput(&lumberjack.Logger{
	// 		Filename:   config.Filename,
	// 		MaxSize:    config.MaxSize,
	// 		MaxBackups: config.MaxBackups,
	// 		MaxAge:     config.MaxAge,
	// 		Compress:   config.Compress,
	// 	})
	// }
	return logger
}

//func InitDefaultLogger() {
//	logrus.SetOutput(os.Stdout)
//	logrus.SetReportCaller(true)
//	logrus.SetFormatter(&logFormatter{})
//	logrus.SetLevel(logrus.DebugLevel)
//}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type logFormatter struct{}

func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var LevelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		LevelColor = gray
	case logrus.WarnLevel:
		LevelColor = yellow
	case logrus.ErrorLevel:
		LevelColor = red
	default:
		LevelColor = blue
	}
	b := &bytes.Buffer{}
	if entry.Buffer != nil {
		b = entry.Buffer
	}
	timeFormat := entry.Time.Format(time.DateTime)
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timeFormat, LevelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m  %s\n", timeFormat, LevelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func ParesLog(config *config.Logger) *logrus.Logger {
	logrusConfig := new(LogrusConfig)
	logrusConfig.LevelStr = config.Level
	var err error
	if logrusConfig.Level, err = logrus.ParseLevel(logrusConfig.LevelStr); err != nil {
		logrus.Warnf("logrusConfig.LevelStr err: %s", err.Error())
	}
	logrusConfig.LogInConsole = config.LogInConsole
	return InitLogrus(logrusConfig)
}
