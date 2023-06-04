package router

import (
	"gin-chat-uwu/middlewares"
	"gin-chat-uwu/services"

	"github.com/gin-gonic/gin"
)

func Chat(ctx *gin.Context) {
	services.Chat(ctx.Writer, ctx.Request)
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/login", services.Login)
	router.POST("/register", services.Register)
	user := router.Group("user")
	{
		user.GET("/chat", Chat)
		user.GET("/msg", services.GetMsg)
		user.DELETE("/", middlewares.JWY(), services.DeleUser)
		user.PUT("/", middlewares.JWY(), services.Update)
	}

	return router

}
