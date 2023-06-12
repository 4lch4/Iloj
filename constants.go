package iloj

import (
	"github.com/4lch4/glogger"
)

var (
	appName  = "Iloj"
	logLevel = 0
)

var logger = glogger.NewLogger(&glogger.NewLoggerInput{
	AppName:  &appName,
	LogLevel: &logLevel,
})

// #region Exported Defaults

// A map of log levels with the key as a string of the log level name or the log level number.
// This means that the following are valid keys:
//
// - "DEBUG" | "0"
//
// - "INFO" | "1"
//
// - "WARNING" | "2"
//
// - "ERROR" | "3"
//
// - "FATAL" | "4"
var LogLevels = map[string]string{
	"DEBUG":   "0",
	"INFO":    "1",
	"WARNING": "2",
	"ERROR":   "3",
	"FATAL":   "4",
	"0":       "DEBUG",
	"1":       "INFO",
	"2":       "WARNING",
	"3":       "ERROR",
	"4":       "FATAL",
}

// #endregion Exported Defaults
