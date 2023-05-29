package models

import (
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	Username string
	Name     string
	CreateAt time.Time
	Data     string
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
