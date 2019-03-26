package main

import (
	"bufio"
	"log"
	"net"
	"net/textproto"
	"os"

	logger "github.com/sirupsen/logrus"

	"github.com/marahin/goredis/protocol"
	"github.com/marahin/goredis/server"
)

func init() {
	switch os.Getenv("ENVIRONMENT") {
	case "development":
		logger.SetLevel(logger.DebugLevel)
	default:
		logger.SetLevel(logger.InfoLevel)
	}

	server.Start()
}

func main() {
	logger.Info("goredis up and running!")

	for {
		select {

		// Accept new clients
		//
		case conn := <-server.NewConnections:
			logger.WithFields(logger.Fields{"clientId": server.ClientCount}).Info("Client connected")
			server.AllClients[conn] = server.ClientCount
			server.ClientCount += 1

			// Receive messages
			go func(conn net.Conn, clientId int) {
				reader := bufio.NewReader(conn)
				tp := textproto.NewReader(reader)
				for {
					incoming, err := tp.ReadLine()
					if err != nil {
						break
					}
					msg := server.Message{ClientId: clientId, Payload: incoming}
					logger.WithFields(logger.Fields{"message": msg}).Info("New message received")
					server.Messages <- msg
				}
				// kill connection on error
				server.DeadConnections <- conn

			}(conn, server.AllClients[conn])

		// Accept messages from connected clients
		//
		case message := <-server.Messages:
			logger.WithFields(logger.Fields{"message": message}).Info("Handling message")
			conn := server.AllClientsInversed()[message.ClientId]

			go func(conn net.Conn, resp []byte) {
				writer := bufio.NewWriter(conn)
				tp := textproto.NewWriter(writer)

				err := tp.PrintfLine(string(resp))

				// If there was an error communicating
				// with them, the connection is dead.
				if err != nil {
					server.DeadConnections <- conn
				}
			}(conn, protocol.Determine(message.Payload).Evaluate().Response)

		// Remove dead clients
		case conn := <-server.DeadConnections:
			logger.WithFields(logger.Fields{"clientId": server.AllClients[conn]}).Info("Client disconnected")
			log.Printf("Client %d disconnected", server.AllClients[conn])
			delete(server.AllClients, conn)
		}
	}

}
