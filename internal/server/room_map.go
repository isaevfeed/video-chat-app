package server

import (
	"fmt"
	"sync"
	"video-chat-app/internal/hash"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomId string) []Participant {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return r.Map[roomId]
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	roomId := hash.HashGenerate(8)
	r.Map[roomId] = []Participant{}

	return roomId
}

func (r *RoomMap) InsertIntoRoom(roomId string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	fmt.Println("Inserting into rom with roomId: ", roomId)
	r.Map[roomId] = append(r.Map[roomId], p)
}

func (r *RoomMap) DeleteRoom(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomId)
}
