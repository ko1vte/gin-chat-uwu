package services

import (
	"gin-chat-uwu/dao"
	"gin-chat-uwu/global"
	"gin-chat-uwu/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func Register(ctx *gin.Context) {
	NewUser := new(models.User)
	NewUser.Name = ctx.PostForm("name")
	NewUser.Username = ctx.PostForm("username")
	NewUser.Password = global.Sha256Encode(ctx.PostForm("password"))
	NewUser.CreatedAt = mysql.NullTime{Time: time.Now(), Valid: true}
	data, err := dao.SelectUserByKEY("username", NewUser.Username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    -1,
			"message": "注册失败，服务器发生错误！",
		})
		return
	}
	log.Println(data)
	if data != nil {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "注册失败，用户名已存在",
		})
		return
	}
	err = dao.AddUser(NewUser)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    "-1",
			"message": "注册失败，服务器发生错误！",
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    "1",
			"message": "注册成功！",
		})
	}

}
