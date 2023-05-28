package router

import (
	"gin-chat-uwu/middlewares"
	"gin-chat-uwu/services"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", services.Login)
	user := router.Group("user")
	{
		user.GET("/login", middlewares.JWY(), services.Login)
		user.POST("/register", middlewares.JWY(), services.Register)
	}

	return router

}
