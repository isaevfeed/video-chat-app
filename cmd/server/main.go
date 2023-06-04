package main

import (
	"fmt"
	"log"
	"net/http"
	"video-chat-app/internal/server"
)

func main() {
	server.AllRooms.Init()

	http.HandleFunc("/create", server.CreateRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	fmt.Println("Starting on the 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("%s", err)
	}
}
