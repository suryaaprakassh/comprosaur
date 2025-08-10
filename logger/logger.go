package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

//TODO: single threaded impl 
// make it thread safe
func GetLogger() *slog.Logger {
	if logger != nil {
		return logger
	}
	file, err := os.Create("comp.log")
	if err != nil {
		panic(err.Error())
	}
	handler := slog.NewTextHandler(file,nil)
	logger = slog.New(handler)
	return logger
}
