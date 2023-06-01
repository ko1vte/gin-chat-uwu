package services

import (
	"gin-chat-uwu/dao"

	"github.com/gin-gonic/gin"
)

func DeleUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	err := dao.DeleUser(username)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "注销失败,出现了一些故障。",
		})
	}
	ctx.JSON(200, gin.H{
		"message": "注销成功",
	})
}
