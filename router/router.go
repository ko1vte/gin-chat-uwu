package router

import (
	"gin-chat-uwu/services"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", services.Login)
	router.POST("/register", services.Register)
	return router
}
