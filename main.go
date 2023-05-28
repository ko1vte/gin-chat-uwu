package main

import (
	"gin-chat-uwu/database"
)

func main() {
	database.InitMysqlDB()
	database.InitRedisdb()
}
