package client

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/BladekTech/blaze/internal/blaze/util"
)

type Client struct {
	addr string
	conn *conn
}

type conn struct {
	inner *net.TCPConn
}

func NewClient(ip string, port int16) Client {
	return Client{
		addr: fmt.Sprint(ip, ":", port),
	}
}

func (client Client) connect() conn {
	addr, err := net.ResolveTCPAddr("tcp", client.addr)
	if err != nil {
		log.Println("ResolveTCPAddr Error:", err)
		os.Exit(1)
	}

	connection, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("DialTCP Error:", err)
		os.Exit(1)
	}

	return conn{
		inner: connection,
	}
}

func (connection conn) read() string {
	buffer := make([]byte, 1024)
	n, err := connection.inner.Read(buffer)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	str := strings.Trim(string(buffer[:n]), "\x00")
	return str
}

func (connection conn) write(data string) {
	connection.inner.Write(util.StrToByteSlice(data + "\n"))
}

func (client Client) Exists(key string) bool {
	conn := client.connect()
	conn.write("exists\n" + key)
	response := conn.read()
	if !util.StartsWith(response, "+") {
		strings.Replace(response, "+", "", 1)
		return util.StartsWith(response, "y")
	} else {
		log.Println("Error:", response)
		return false
	}
}

func (client Client) Get(key string) string {
	conn := client.connect()
	conn.write("get\n" + key)
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error:", response)
		return ""
	} else {
		return strings.Replace(response, "+", "", 1)
	}
}

func (client Client) Set(key string, value string) {
	conn := client.connect()
	conn.write("set\n" + key + "\n" + value)
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error:", response)
	}
}

func (client Client) Ping() {
	conn := client.connect()
	conn.write("ping")
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error:", response)
	} else {
		log.Println("Pong!")
	}
}

func (client Client) Clear() {
	conn := client.connect()
	conn.write("clear")
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error:", response)
	}
}

func (client Client) Delete(key string) {
	conn := client.connect()
	if client.Exists(key) {
		conn.write("delete\n" + key)
		response := conn.read()
		if !util.StartsWith(response, "+") {
			log.Println("Error:", response)
		}
	}
}

func (client Client) Update(key string, value string) {
	conn := client.connect()
	conn.write("update\n" + key + "\n" + value)
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error:", response)
	}
}
