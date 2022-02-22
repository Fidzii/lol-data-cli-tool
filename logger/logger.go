package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	infoLoggerPrefixStyle    string = color.New(color.FgBlue).Add(color.Bold).Sprint("[INFO]")
	warningLoggerPrefixStyle string = color.New(color.FgYellow).Add(color.Bold).Sprint("[WARNING]")
	errorLoggerPrefixStyle   string = color.New(color.FgRed).Add(color.Bold).Sprint("[ERROR]")
)

var (
	infoLogger    *log.Logger = log.New(os.Stdout, getPrefix(infoLoggerPrefixStyle), log.Ltime)
	warningLogger *log.Logger = log.New(os.Stdout, getPrefix(warningLoggerPrefixStyle), log.Ltime)
	errorLogger   *log.Logger = log.New(os.Stdout, getPrefix(errorLoggerPrefixStyle), log.Ltime)
)

func getPrefix(s string) string {
	return s + " "
}

func Info(s string) {
	infoLogger.Println(s)
}

func Warning(s string) {
	warningLogger.Println(s)
}

func Error(s string) {
	errorLogger.Println(s)
}
