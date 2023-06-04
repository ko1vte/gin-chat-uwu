package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-chat-uwu/database"
	"gin-chat-uwu/models"
	"log"
)

func SaveMsg(msg []byte) error {
	client, err := database.InitRedisdb()
	defer client.Close()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	err = client.LPush(context.Background(), "msg", msg).Err()
	if err != nil {
		return err
	}
	log.Println("redis存入成功", string(msg))
	return nil
}

func GetAllMsg(key string) ([]models.Message, error) {
	// 获取列表中的所有元素
	client, err := database.InitRedisdb()
	defer client.Close()
	if err != nil {
		return nil, err
	}
	result, err := client.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	// 解析JSON字符串为消息对象
	var messages []models.Message
	for _, msgJSON := range result {
		var message models.Message
		err := json.Unmarshal([]byte(msgJSON), &message)
		if err != nil {
			fmt.Println("无法解析消息：", err)
			continue
		}
		messages = append(messages, message)
	}

	return messages, nil
}
