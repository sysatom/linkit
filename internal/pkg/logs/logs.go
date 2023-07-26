package logs

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/sysatom/linkit/internal/pkg/constant"
	"github.com/sysatom/linkit/internal/pkg/setting"
	"io"
	"log"
	"os"
)

var l zerolog.Logger

func Init() {
	var writer []io.Writer
	// log file
	logFileName := fmt.Sprintf("%s.log", constant.AppId)
	logPath := setting.Get().LogPath
	if _, err := os.Stat(logPath); !os.IsNotExist(err) {
		logFilePath := fmt.Sprintf("%s/%s", logPath, logFileName)
		var logFile *os.File
		_, err = os.Stat(logFilePath)
		if os.IsNotExist(err) {
			logFile, err = os.Create(logFilePath)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			logFile, err = os.OpenFile(logFilePath, os.O_RDWR|os.O_APPEND, 0666)
		}
		console := zerolog.ConsoleWriter{Out: logFile, NoColor: true, TimeFormat: zerolog.TimeFieldFormat}
		writer = append(writer, console)
	}
	// console
	console := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat}
	writer = append(writer, console)

	multi := zerolog.MultiLevelWriter(writer...)
	l = zerolog.New(multi).With().Timestamp().Logger()
}

func Debug(format string, a ...any) {
	l.Debug().Caller(1).Msg(fmt.Sprintf(format, a...))
}

func Info(format string, a ...any) {
	l.Info().Caller(1).Msg(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	l.Warn().Caller(1).Msg(fmt.Sprintf(format, a...))
}

func Error(err error) {
	l.Error().Caller(1).Msg(err.Error())
}

func Fatal(format string, a ...any) {
	l.Fatal().Caller(1).Msg(fmt.Sprintf(format, a...))
}

func Panic(format string, a ...any) {
	l.Panic().Caller(1).Msg(fmt.Sprintf(format, a...))
}
