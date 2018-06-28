package tools

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var loggerJSON *os.File

// Log configures the log file.
func Log() *log.Logger {
	if logger == nil {
		if f, err := os.OpenFile("go-cli.log", os.O_RDWR|os.O_CREATE, 0666); err == nil {
			logger = log.New(f, "GOCLI: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
		} else {
			panic("log file could not be opened")
		}
		logger.Println("***** LOGGER HAS BEEN INITIALIZED *****")
	}
	return logger
}

func inlog(format string, params ...interface{}) {
	Log().Printf(format, params...)
}

// Tracer logs the trace..
func Tracer(format string, params ...interface{}) {
	//Log().Printf(fmt.Sprintf("[TRACER]:%s", format), params...)
	inlog(fmt.Sprintf("[TRACER]:%s", format), params...)
}

// Tester logs the test log..
func Tester(format string, params ...interface{}) {
	//Log().Printf(fmt.Sprintf("[TESTER]:%s", format), params...)
	inlog(fmt.Sprintf("[TESTER]:%s", format), params...)
}

// LogJSON configures the JSON log file.
func LogJSON() *os.File {
	if loggerJSON == nil {
		var err error
		if loggerJSON, err = os.OpenFile("go-cli-json.log", os.O_RDWR|os.O_CREATE, 0666); err != nil {
			panic(err)
		}
		logger.Println("***** LOGGER-JSON HAS BEEN INITIALIZED *****")
	}
	return loggerJSON
}
