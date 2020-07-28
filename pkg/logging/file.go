package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log_"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return LogSavePath
}

func getLogFileFullPath() string {
	return fmt.Sprintf("%s%s%s.%s",
		getLogFilePath(),
		LogSaveName,
		time.Now().Format(TimeFormat),
		LogFileExt)
}

func openLogFile(filepath string) *os.File {
	if _, err := os.Stat(filepath); err != nil {
		switch {
		case os.IsNotExist(err):
			mkDir()
		case os.IsPermission(err):
			log.Fatalf("permission denied: %v\n", err)

		}
	}

	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("fail to open file: %v\n", err)
	}

	return fd
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
