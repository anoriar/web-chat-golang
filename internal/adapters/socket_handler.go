package adapters

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // оставим без изменений
var connections = []*websocket.Conn{}

func removeConn(slice []*websocket.Conn, val *websocket.Conn) []*websocket.Conn {
	index := -1
	for i, v := range slice {
		if v == val {
			index = i
			break
		}
	}

	if index != -1 {
		if index < len(slice)-1 {
			copy(slice[index:], slice[index+1:])
		}
		slice = slice[:len(slice)-1]
	}

	return slice
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	connections = append(connections, conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			connections = removeConn(connections, conn)
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Server: %s", message)
		for _, c := range connections {
			if c != conn {
				err = c.WriteMessage(messageType, message)
				if err != nil {
					connections = removeConn(connections, conn)
					log.Println("Error during message writing:", err)
					break
				}
			}
		}
	}
	connections = removeConn(connections, conn)
}
