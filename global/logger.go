package global

import (
	"log"
	"os"
)

var Logger *log.Logger

//初始化mysql的日志
func InitLog() {
	file, err := os.OpenFile("sql.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("创建日志文件 \"sql.log\"失败")
	}
	Logger = log.New(file, "mysql", log.Default().Flags())
}
