package database

import (
	"context"
	"database/sql"
	"fmt"
	"gin-chat-uwu/global"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
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
	log.Println("mysql connecting successful!")
	return db, err
}

func InitRedisdb() (*redis.Client, error) {
	ctx := context.Background()
	config, err := global.ReadConfig("config/config.yaml")
	if err != nil {
		return nil, err
	}

	// 根据配置信息创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	// 关闭 Redis 连接
	log.Println("Redis connecting successful!")
	return client, err
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(100) NOT NULL,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			name VARCHAR(100) NOT NULL
		)
	`)
	if err != nil {
		log.Println("初始化表失败")
		return err
	}
	return nil
}
