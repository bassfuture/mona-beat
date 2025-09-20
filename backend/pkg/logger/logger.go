package logger

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func Init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(v ...interface{}) {
	if InfoLogger != nil {
		InfoLogger.Println(v...)
	}
}

func Warning(v ...interface{}) {
	if WarningLogger != nil {
		WarningLogger.Println(v...)
	}
}

func Error(v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Println(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if InfoLogger != nil {
		InfoLogger.Printf(format, v...)
	}
}

func Warningf(format string, v ...interface{}) {
	if WarningLogger != nil {
		WarningLogger.Printf(format, v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if ErrorLogger != nil {
		ErrorLogger.Printf(format, v...)
	}
}