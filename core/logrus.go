package core

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"blog/config"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := gray
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.InfoLevel:
		levelColor = blue
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	b := &bytes.Buffer{}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")

	if entry.HasCaller() {
		file := path.Base(entry.Caller.File)
		line := entry.Caller.Line
		function := entry.Caller.Function
		fmt.Fprintf(b, "\x1b[%dm[%s][%s] [%s] %s:%d %s: %s\x1b[0m\n",
			levelColor,
			config.Config.Logger.Prefix,
			timestamp,
			switchLevel(entry.Level),
			file,
			line,
			function,
			entry.Message,
		)
	} else {
		fmt.Fprintf(b, "\x1b[%dm[%s][%s] [%s] %s\x1b[0m\n",
			levelColor,
			config.Config.Logger.Prefix,
			timestamp,
			switchLevel(entry.Level),
			entry.Message,
		)
	}

	return b.Bytes(), nil
}

func switchLevel(level logrus.Level) string {
	return level.String()
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(config.Config.Logger.ShowLine)
	mLog.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)

	InitDefaultLogger()
	return mLog
}

func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(config.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
