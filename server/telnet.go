package server

import (
	"fmt"
	"net"
	"os"

	_ "github.com/marahin/goredis/config"
)

type Message struct {
	ClientId int
	Payload  string
}

var NewConnections = make(chan net.Conn)
var DeadConnections = make(chan net.Conn)
var Messages = make(chan Message)
var AllClients = map[net.Conn]int{}
var AllClientsInversed = func() map[int]net.Conn {
	n := make(map[int]net.Conn)
	for k, v := range AllClients {
		n[v] = k
	}

	return n
}
var ClientCount = 0
var Host = func() string {
	var str string

	if str = os.Getenv("HOST"); str == "" {
		str = "0.0.0.0"
	}

	return str
}()
var Port = func() string {
	var str string

	if str = os.Getenv("PORT"); str == "" {
		str = "6379"
	}

	return str
}()
var Instance = func() net.Listener {
	ins, err := net.Listen("tcp", Dsn)

	if err != nil {
		panic(err)
	}

	return ins
}()
var Dsn = Host + ":" + Port

func Start() {
	go func() {
		for {
			conn, err := Instance.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			NewConnections <- conn
		}
	}()
}
