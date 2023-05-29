package services

import (
	"fmt"
	"gin-chat-uwu/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

func wsChat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	username := query.Get("username")
	var isvalida = true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrader(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	node := &models.Node{
		Conn:     conn,
		Username: username,
		Data:     make(chan []byte, 50),
	}

	models.RwLocker.Lock()
	models.ClientMap[username] = node
	models.RwLocker.Unlock()

	go sendMessage(node)

}

func sendMessage(node *models.Node) {
	for {
		select {
		case data := <-node.Data:
			_, err := node.Conn.Write(websocket.TextMessage, data)
			if err != nil {
				log.Println("写入消息失败", err)
			}
		}
	}
}
