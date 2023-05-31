package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string    `json:"username"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"time"`
	Data     string    `json:"data"`
}

func (m *Message) MsgTableName() string {
	return "message"
}

type Node struct {
	Conn     *websocket.Conn
	Username string
	Data     chan []byte
}

var ClientMap map[string]*Node = make(map[string]*Node, 0)
var RwLocker sync.RWMutex
