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
	router.POST("/login", services.Login)
	user := router.Group("user")
	{
		user.GET("/login", middlewares.JWY(), services.Login)
		user.POST("/register", middlewares.JWY(), services.Register)
		user.GET("/chat", Chat)
		user.GET("/dele", middlewares.JWY(), services.DeleUser)
	}

	return router

}
