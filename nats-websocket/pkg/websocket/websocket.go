package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Websocket struct {
	Clients   map[*websocket.Conn]bool
	Upgrader  websocket.Upgrader
	Mutex     *sync.Mutex
	Broadcast chan []byte
}

func StartWebsocketServer() *Websocket {
	ws := &Websocket{
		Mutex:     &sync.Mutex{},
		Clients:   make(map[*websocket.Conn]bool),
		Upgrader:  websocket.Upgrader{},
		Broadcast: make(chan []byte),
	}

	return ws
}

func (ws *Websocket) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ws.Mutex.Lock()

	// Register the new client.
	ws.Clients[conn] = true
	ws.Mutex.Unlock()

	log.Println("Total Registered Clients: ", len(ws.Clients))
	// Read messages from client in order
	// to handle close connection by client.
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			switch true {
			case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway):
				log.Printf("Connection for client %s close", conn.RemoteAddr())
			case websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure):
				log.Printf("Connection for client %s close abnormally, Error: %s", conn.RemoteAddr(), err.Error())
			default:
				log.Printf("Unhandle error: %s", err.Error())
			}

			// Remove connection from the active viewers and break the loop
			ws.Mutex.Lock()
			delete(ws.Clients, conn)
			ws.Mutex.Unlock()
			break
		}
	}
}

func (ws *Websocket) BroadcastMessageToClients() {
	for {
		select {
		case msg := <-ws.Broadcast:
			for client := range ws.Clients {
				err := client.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Println("Error writing message:", err)
					client.Close()
					ws.Mutex.Lock()
					delete(ws.Clients, client)
					ws.Mutex.Unlock()
				}
			}
		}
	}
}
