package global

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger() {
	Logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("初始化日志失败：", err.Error())
	}
	zap.RedirectStdLog(Logger)
}
