package client

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/BladekTech/Blaze/internal/blaze/util"
)

type Client struct {
	addr string
	conn *Conn
}

type Conn struct {
	inner *net.TCPConn
}

func NewClient(ip string, port int16) Client {
	return Client{
		addr: fmt.Sprint(ip, ":", port),
	}
}

func (client Client) connect() Conn {
	addr, err := net.ResolveTCPAddr("tcp", client.addr)
	if err != nil {
		log.Println("ResolveTCPAddr Error:", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("DialTCP Error:", err)
		os.Exit(1)
	}

	return Conn{
		inner: conn,
	}
}

func (conn Conn) read() string {
	buffer := make([]byte, 1024)
	n, err := conn.inner.Read(buffer)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	str := strings.Trim(string(buffer[:n]), "\x00")
	return str
}

func (conn Conn) write(data string) {
	conn.inner.Write(util.StrToByteSlice(data + "\n"))
}

func (client Client) Get(key string) string {
	conn := client.connect()
	conn.write("get\n" + key)
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error: ", response, util.StrToByteSlice(response))
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
		log.Println("Error: ", response, util.StrToByteSlice(response))
	}
}

func (client Client) Ping() {
	conn := client.connect()
	conn.write("ping")
	response := conn.read()
	if !util.StartsWith(response, "+") {
		log.Println("Error: ", response, util.StrToByteSlice(response))
	} else {
		log.Println("Pong!")
	}
}
