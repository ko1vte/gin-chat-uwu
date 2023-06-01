package services

import (
	"encoding/json"
	"fmt"
	"gin-chat-uwu/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	username := query.Get("username")
	var isvalida = true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)
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

	go SendMessage(node)
	go reccMessage(node)

}

func SendMessage(node *models.Node) {
	for {
		select {
		case data := <-node.Data:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("写入消息失败", err)
			}
		}
	}
}

func reccMessage(node *models.Node) {
reconve:
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println("读取消息失败", err)
			goto reconve
		}
		msg := models.Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Println("json解析失败", err)
			goto reconve
		}
		tarnode, ok := models.ClientMap[msg.Username]
		if !ok {
			log.Println("没有node", msg.Username)
			goto reconve
		}
		tarnode.Data <- data
		log.Println("发送成功", string(data))
	}
}

// {"Username" :"su15vte",
// "Name" :"su15vte",
// "CreateAt" :"2002-10-15",
// "Data":     "Hello World"}
