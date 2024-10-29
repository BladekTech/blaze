package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/BladekTech/blaze/internal/blaze/store"
	"github.com/BladekTech/blaze/pkg/protocol"
)

type Listener struct {
	inner net.Listener
}

type Conn struct {
	inner net.Conn
}

func (listener Listener) Accept() Conn {
	conn, err := listener.inner.Accept()
	if err != nil {
		log.Println("Error accepting connection: ", err.Error())
	}

	return Conn{
		inner: conn,
	}
}

func (conn Conn) HandleClient(store store.Store) {
	handleClient(conn.inner, store)
}

func StartTcpServer(port int16) Listener {
	l, err := net.Listen("tcp", "0.0.0.0:"+fmt.Sprint(port))
	if err != nil {
		log.Println("Failed to bind to port", fmt.Sprint(port))
		os.Exit(1)
	}

	log.Println("blaze running on 0.0.0.0:" + fmt.Sprint(port))

	return Listener{
		inner: l,
	}
}

func handleClient(conn net.Conn, store store.Store) {
	log.Println("Connected with", conn.RemoteAddr().String())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	str := string(buffer[:n])
	scmd := strings.Split(str, "\n")
	scmd = scmd[:len(scmd)-1]

	log.Printf("From %s: %s\n", conn.RemoteAddr().String(), scmd)

	cmd := protocol.Command{
		Name: strings.Trim(strings.ToLower(scmd[0]), "\r\n\x00"),
		Args: scmd[1:],
	}
	result := processCommand(cmd, store)

	conn.Write(result.Data.ToBytes())
	conn.Close()
}

func processCommand(command protocol.Command, store store.Store) protocol.Result {
	switch command.Name {
	case protocol.CMD_GET:
		result := store.Get(command.Args[0])
		if result.Status != protocol.STATUS_OK {
			log.Println("Error:", result.Status)
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("-get\n"),
			}
		} else {
			log.Println(fmt.Sprint("+", *result.Result))
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData(fmt.Sprint("+", *result.Result)),
			}
		}
	case protocol.CMD_SET:
		result := store.Set(command.Args[0], command.Args[1])
		if result.Status != protocol.STATUS_OK {
			log.Println("Error:", result.Status)
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("-set\n"),
			}
		} else {
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("+set\n"),
			}
		}
	case protocol.CMD_DELETE:
		result := store.Delete(command.Args[0])
		if result.Status != protocol.STATUS_OK {
			log.Println("Error:", result.Status)
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("-delete\n"),
			}
		} else {
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("+delete\n"),
			}
		}
	case protocol.CMD_CLEAR:
		result := store.Clear()
		if result.Status != protocol.STATUS_OK {
			log.Println("Error:", result.Status)
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("-clear\n"),
			}
		} else {
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("+clear\n"),
			}
		}
	case protocol.CMD_UPDATE:
		result := store.Update(command.Args[0], command.Args[1])
		if result.Status != protocol.STATUS_OK {
			log.Println("Error:", result.Status)
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("-update\n"),
			}
		} else {
			return protocol.Result{
				Status: result.Status,
				Error:  nil,
				Data:   protocol.NewData("+update\n"),
			}
		}
	case protocol.CMD_EXISTS:
		exists := store.Exists(command.Args[0])
		if exists {
			return protocol.Result{
				Status: 200,
				Error:  nil,
				Data:   protocol.NewData("+y\n"),
			}
		} else {
			return protocol.Result{
				Status: 200,
				Error:  nil,
				Data:   protocol.NewData("+n\n"),
			}
		}
	case protocol.CMD_PING:
		return protocol.Result{
			Status: protocol.STATUS_OK,
			Error:  nil,
			Data:   protocol.NewData("+pong\n"),
		}
	default:
		return protocol.Result{
			Status: protocol.STATUS_NO_SUCH_COMMAND,
			Error:  nil,
			Data:   protocol.NewData("-noc\n"),
		}
	}

}
