package services

import (
	"gin-chat-uwu/dao"
	"gin-chat-uwu/global"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	data, err := dao.SelectUserByKEY("username", username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    -1,
			"message": "登陆失败，服务端发生了一些错误",
		})
	}
	if data == nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "用户不存在",
		})
	}
	if username != data.Username || global.Sha256Encode(password) != data.Password {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "登陆失败，帐号或密码错误",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":     1,
			"message":  "登陆成功",
			"username": username,
		})
	}
}
