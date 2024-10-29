package main

import (
	"log"

	blaze "github.com/BladekTech/blaze/pkg/client"
	"github.com/BladekTech/blaze/pkg/protocol"
)

func main() {
	client := blaze.NewClient("localhost", protocol.DEFAULT_PORT)
	client.Set("key", "value")
	result := client.Get("key")
	log.Println(result)
	client.Clear()
	result2 := client.Get("key")
	log.Println(result2)
}
