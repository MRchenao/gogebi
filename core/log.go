package core

import (
	"gebi/app/Http/Serializer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

func initLog() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(logFile(makeFilePath()))
	log.SetLevel(log.WarnLevel)
}

func logFile(logpath string) *os.File {
	if file, err := os.OpenFile(path.Join(logpath, "log.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666); err == nil {
		return file
	} else {
		Serializer.Err(2342, "日志文件创建失败", err)
	}
	return nil
}

func makeFilePath() string {
	logpath := viper.GetString("log.logpath")
	_, exist := os.Stat(logpath)
	if os.IsNotExist(exist) {
		err := os.MkdirAll(logpath, 0666)
		if err != nil {
			Serializer.Err(2342, "日志文件夹创建失败", err)
		}
	}
	return logpath
}
