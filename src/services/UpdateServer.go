package services

import (
	"gin-chat-uwu/dao"
	"log"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.PostForm("password")
	err := dao.UpdatePwd(password, username)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "修改密码失败",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "修改密码成功",
		"username": username,
	})
}
