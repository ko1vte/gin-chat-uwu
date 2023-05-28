package main

import "gin-chat-uwu/router"

func main() {
	r := router.InitRouter()
	r.Run()
}
