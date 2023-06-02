package services

import (
	"gin-chat-uwu/dao"
	"log"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	username := ctx.Query("username")
	name := ctx.PostForm("name")
	err := dao.Updatename(name, username)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "修改昵称失败",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "修改昵称成功",
		"username": username,
		"name":     name,
	})
}
