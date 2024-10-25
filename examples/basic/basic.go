package main

import blaze "github.com/BladekTech/blaze/pkg/client"

func main() {
	// initialize the blaze client with host and port
	client := blaze.NewClient("localhost", 6589)

	// Pings the server
	// Server should response with +pong.
	// This method prints "Pong!" if the server responds correctly
	client.Ping()
}
