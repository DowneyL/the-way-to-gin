package logging

import (
	"fmt"
	"github.com/DowneyL/the-way-to-gin/pkg/file"
	"github.com/DowneyL/the-way-to-gin/pkg/setting"
	"os"
	"time"
)

func getLogFilePath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.LogSavePath
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExtend)
}

func openLogFile(filename, filepath string) (*os.File, error) {
	dir, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v\n", err)
	}

	src := dir + "/" + filepath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("permission denied src: %v\n", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, err
	}

	return file.Open(src+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
