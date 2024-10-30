# Blaze

<div align="center">

<img src="assets/png/icon-with-name.png" width="256px" alt="Blaze Logo">

A fast, temporary key-value store

</div>

## Table of Contents

- [Introduction](#introduction)
- [Usage](#usage)
  - [Installation](#installation)
    - [OS Support](#os-support)
      - [Windows](#on-windows)
      - [Linux](#on-linux)
      - [Mac](#on-mac)
  - [Client API](#client-api)
  - [Blaze CLI](#blaze-cli)
- [Building](#building)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Blaze is an alternative to [redis](https://redis.io). It does not use the RESP protocol, and instead uses a custom protocol. While the protocols have some similarities, Blaze is much simpler in its implementation and specification.

Blaze is written with Go 1.22.4

## Usage

### Installation

To get started, install the blaze server executable by downloading it from the [Releases](https://github.com/BladekTech/blaze/releases/latest).

Make sure you install the correct version for your operating system and machine.

> Note that Mac is labeled `darwin`

For example, a computer running Linux

#### OS Support

There are Windows, Linux, and Mac builds available. Please note, **I have only tested `blaze-server-windows-amd64.exe`**.

##### On windows

You can either run the executable normally (which opens a command prompt/terminal) with the server running, or use the `start blaze-server.exe` command, then press `Ctrl + C` which will "start the app headlessly".

##### On Linux

The executable file provided is untested. Please open an issue if you encounter problems.

##### On Mac

As with Linux, he executable file provided is untested. Please open an issue if you encounter problems.

Also, please keep in mind that I am not familiar with Mac. If you have any suggestions to improve the build for Mac, please open an issue.

### Client API

To use the Go client first run `go get github.com/BladekTech/blaze`

The following showcases all of the functionality (so far)!

```go
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

```

### Blaze CLI

The Blaze CLI is a work in progress, it will likely live in another GitHub repository in the future. For now, it serves the purpose of a simple testing file.

## Building

There is a `Makefile` included. It supports Windows, Mac, and Linux (note that not all Linux arches are supported by default).
If your arch is unsupported/unlisted, run `GOOS=<YOUR OS HERE> GOARCH=<YOUR ARCH HERE> go build ./cmd/blaze-server`. Your executable is `./blaze-server(.exe)`.

## Contributing

Blaze is not currently accepting contributions.

## License

[MIT](https://github.com/BladekTech/blaze/blob/main/LICENSE)
