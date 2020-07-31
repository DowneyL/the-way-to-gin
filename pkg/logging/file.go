package logging

import (
	"fmt"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"log"
	"os"
	"time"
)

func getLogFilePath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
}

func getLogFileFullPath() string {
	return fmt.Sprintf("%s%s%s.%s",
		getLogFilePath(),
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExtend)
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
