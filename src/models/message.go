package models

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string    `json:"username"`
	Msg      string    `json:"data"`
	Time     time.Time `json:"time"`
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
