package main

import (
	"log"

	blaze "github.com/BladekTech/blaze/pkg/client"
	"github.com/BladekTech/blaze/pkg/protocol"
)

func main() {
	client := blaze.NewClient("localhost", protocol.DEFAULT_PORT)
	// client.Set("key", "value")
	value := client.Get("key")
	log.Println("key", value)
}
