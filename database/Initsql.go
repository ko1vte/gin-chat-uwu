package database

import (
	"database/sql"
	"fmt"
	"github/su15vte/gin-chat-uwu/global"
	"log"
)

func InitMysqlDB() (*sql.DB, error) {
	config, err := global.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatal("Reading config file error", err)
		return nil, err
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
		return nil, err
	}
}
