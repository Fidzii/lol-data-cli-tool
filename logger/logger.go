package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	infoLoggerPrefixStyle    string = color.New(color.FgBlue).Add(color.Bold).Sprint("[INFO]")
	warningLoggerPrefixStyle string = color.New(color.FgYellow).Add(color.Bold).Sprint("[WARNING]")
	errorLoggerPrefixStyle   string = color.New(color.FgRed).Add(color.Bold).Sprint("[ ERROR ]")
	fatalLoggerPrefixStyle   string = color.New(color.BgRed).Add(color.Bold).Sprint("[[ FATAL ]]")
)

var (
	infoLogger    *log.Logger = log.New(os.Stdout, getPrefix(infoLoggerPrefixStyle), log.Ltime)
	warningLogger *log.Logger = log.New(os.Stdout, getPrefix(warningLoggerPrefixStyle), log.Ltime)
	errorLogger   *log.Logger = log.New(os.Stdout, getPrefix(errorLoggerPrefixStyle), log.Ltime)
	fatalLogger   *log.Logger = log.New(os.Stdout, getPrefix(fatalLoggerPrefixStyle), log.Ltime)
)

func getPrefix(s string) string {
	return s + " "
}

func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

func Warning(v ...interface{}) {
	warningLogger.Println(v...)
}

func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

func Fatal(v ...interface{}) {
	fatalLogger.Fatalln(v...)
}
