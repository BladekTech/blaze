package main

import (
	"github.com/BladekTech/Blaze/internal/blaze/server"
	"github.com/BladekTech/Blaze/internal/blaze/store"
)

func main() {
	listener := server.StartTcpServer(6379)
	store := store.NewStore()

	for {
		conn := listener.Accept()
		go conn.HandleClient(store)
	}
}
