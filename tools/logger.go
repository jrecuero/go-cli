package tools

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var logger *log.Logger
var loggerJSON *os.File
var mutex sync.Mutex
var logfile *os.File

// Log configures the log file.
func Log() *log.Logger {
	if logger == nil {
		if logfile, err := os.OpenFile("go-cli.log", os.O_RDWR|os.O_CREATE, 0666); err == nil {
			logger = log.New(logfile, "GOCLI: ", log.Ldate|log.Ltime|log.Lmicroseconds)
		} else {
			panic("log file could not be opened")
		}
		logger.Println("***** LOGGER HAS BEEN INITIALIZED *****")
	}
	return logger
}

// CloseLog the logfile.
func CloseLog() {
	logfile.Close()
}

func inlog(mod string, format string, params ...interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	pc, filename, line, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		methodname := filepath.Base(details.Name())
		//index := strings.LastIndex(methodname, ".")
		//funcname := methodname[index+1:]
		ifname := strings.LastIndex(filename, "/")
		fname := filename[ifname+1:]
		Log().Printf(fmt.Sprintf("[%s] %s:%d %s() ||| %s", mod, fname, line, methodname, format), params...)
	} else {
		Log().Printf(fmt.Sprintf("[%s] ||| %s", mod, format), params...)
	}
}

// Error logs the trace..
func Error(format string, params ...interface{}) {
	inlog("ERROR", format, params...)
}

// Warning logs the trace..
func Warning(format string, params ...interface{}) {
	inlog("WARNING", format, params...)
}

// Info logs the trace..
func Info(format string, params ...interface{}) {
	inlog("INFO", format, params...)
}

// Tracer logs the trace..
func Tracer(format string, params ...interface{}) {
	inlog("TRACER", format, params...)
}

// Debug logs the trace..
func Debug(format string, params ...interface{}) {
	inlog("DEBUG", format, params...)
}

// Tester logs the test log..
func Tester(format string, params ...interface{}) {
	inlog("TERSTER", format, params...)
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

// ToDisplay sends the formated string to the prompt display.
func ToDisplay(format string, params ...interface{}) {
	fmt.Printf(format, params...)
}

// ERROR sends the formated string to the prompt display.
func ERROR(err error, todisp bool, format string, params ...interface{}) error {
	if err == nil {
		err = fmt.Errorf(format, params...)
	}
	Error(format, params...)
	if todisp {
		ToDisplay(format, params...)
	}
	return err
}
