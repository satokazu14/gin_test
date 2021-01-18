package config

import (
	"gin_test/entity"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients map[*websocket.Conn]bool

	Broadcast chan entity.Bid
}

func NewHub() *Hub {
	val := &Hub{
		Broadcast: make(chan entity.Bid),
		Clients:   make(map[*websocket.Conn]bool),
	}
	return val
}
