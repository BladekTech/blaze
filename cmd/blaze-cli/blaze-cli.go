package main

import (
	"log"

	blaze "github.com/BladekTech/Blaze/pkg/client"
)

func main() {
	client := blaze.NewClient("localhost", 6379)
	// client.Set("key", "value")
	value := client.Get("key")
	log.Println("key", value)
}
