package websocket

import (
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
    ws, _ := upgrader.Upgrade(w, r, nil)
    defer ws.Close()

    for {
        var msg map[string]interface{}
        if err := ws.ReadJSON(&msg); err != nil {
            break
        }
        Broadcast(msg)
    }
}

var clients = make(map[*websocket.Conn]bool)

func Broadcast(msg interface{}) {
    for client := range clients {
        if err := client.WriteJSON(msg); err != nil {
            client.Close()
            delete(clients, client)
        }
    }
}