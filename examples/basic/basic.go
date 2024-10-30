package main

import (
	"fmt"

	// This is a choice of style (I am unsure of best practices in Go modules)
	blaze "github.com/BladekTech/blaze/pkg/client"
	"github.com/BladekTech/blaze/pkg/protocol"
)

func main() {
	// initialize the blaze client with host and port
	client := blaze.NewClient("localhost", protocol.DEFAULT_PORT)

	// Pings the server
	// Server should response with +pong.
	// This method prints "Pong!" if the server responds correctly
	client.Ping()

	// Sets "key" to "value"
	client.Set("key", "value")

	// Checks if "key" exists (is a key)
	exists := client.Exists("key")
	if exists {
		// Note that Get will fail if "key" doesn't exist
		// Gets the value of key "key"
		value := client.Get("key")
		if value != "value" {
			fmt.Println("this should never happen")
		}
	}

	// We have to use update when overwriting a key to prevent accidents
	client.Update("key", "value but *different*")

	// This will delete "key" if it exists
	client.Delete("key")

	// Again let's set "key" to "value"
	client.Set("key", "value")

	// Clear simply deletes each key-value pair
	client.Clear()
}
