package services

import (
	"gin-chat-uwu/dao"

	"github.com/gin-gonic/gin"
)

func GetMsg(ctx *gin.Context) {
	messages, err := dao.GetAllMsg("msg")
	if err != nil {
		ctx.JSON(500, gin.H{"error": "无法获取消息"})
		return
	}

	// 返回所有消息
	ctx.JSON(200, gin.H{
		"messages": messages,
	})
}
