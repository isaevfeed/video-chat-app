package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast
		for _, client := range AllRooms.Map[msg.RoomID] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Printf("%s", err)
					client.Conn.Close()
				}
			}
		}
	}
}

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomId := AllRooms.CreateRoom()

	type resp struct {
		RoomId string `json:"roomId"`
	}

	json.NewEncoder(w).Encode(resp{RoomId: roomId})
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomId, ok := r.URL.Query()["roomId"]

	if !ok {
		log.Println("roomId is missing")
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	AllRooms.InsertIntoRoom(roomId[0], false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatalf("%s", err)
		}

		msg.Client = ws
		msg.RoomID = roomId[0]

		broadcast <- msg
	}
}
