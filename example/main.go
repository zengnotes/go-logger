package main

import (
	"github.com/zengnotes/go-logger"
	"os"
)

var (
	Log          *logger.Logger
)

func main() {
	//初始化
	//"FNST", "FINE", "DEBG", "TRAC", "INFO", "WARN", "EROR", "CRIT"
	LogLevel := "DEBG"
	LogLevel_int := logger.StringToLevel(LogLevel)
	LogPath := "/tmp/go-logger"
	os.MkdirAll(LogPath, 0777)
	LogMaxSize := 1024
	fileLog := logger.NewFileLogWriter(LogPath, true)
	fileLog = fileLog.SetRotateSize(LogMaxSize)

	if logger.DEBUG == LogLevel_int {
		Log = &logger.Logger{
			"stdout": &logger.Filter{LogLevel_int, logger.NewConsoleLogWriter()},
		}
	} else {
		Log = &logger.Logger{
			"fileout": &logger.Filter{LogLevel_int, fileLog},
		}
	}

	Log.Log(logger.DEBUG, "sdfsdfsdfsdfsdfsdfdsfsdf", "");

	//等待写完
	defer Log.Close()
}
