package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ccheeliang/nats-example/nats-websocket/pkg/websocket"
	"github.com/nats-io/nats.go"
)

type Server struct {
	websocket *websocket.Websocket
	natsWs    *nats.Conn
}

func StartServer(natsWebsocketUrl string) (*Server, error) {
	nc, err := nats.Connect(natsWebsocketUrl)
	return &Server{
		websocket: websocket.StartWebsocketServer(),
		natsWs:    nc,
	}, err
}

func (server *Server) subscribeSubject(subject string) {
	server.natsWs.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
		server.websocket.Broadcast <- msg.Data
	})
}

func (server *Server) ListenAndServe(port string) {
	log.Printf("Server list and serve at %s\n", port)

	go server.websocket.BroadcastMessageToClients()

	server.subscribeSubject("foo")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	http.HandleFunc("/ws", server.websocket.HandleWebSocket)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
