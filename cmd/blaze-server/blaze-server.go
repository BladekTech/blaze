package main

import (
	"os"

	"github.com/BladekTech/blaze/internal/blaze/server"
	"github.com/BladekTech/blaze/internal/blaze/store"
	"github.com/BladekTech/blaze/internal/blaze/util"
	"github.com/BladekTech/blaze/pkg/protocol"
)

func main() {
	port := protocol.DEFAULT_PORT

	if len(os.Args) > 1 {
		port = int16(util.Atoi(os.Args[1]))
	}

	listener := server.StartTcpServer(port)
	store := store.NewStore()

	for {
		conn := listener.Accept()
		go conn.HandleClient(store)
	}

}
