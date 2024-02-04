package main

import (
	"fmt"

	"github.com/jonavdm/room-notifications/list"
	"github.com/jonavdm/room-notifications/rooms"
)

const ROOM_URL string = ""

func main() {
	r := rooms.GetRooms()
	fmt.Println(len(r))

	for _, room := range r {
		if ok := list.IsInList(room.ID); !ok {
			list.Add(room.ID)
		}
	}
}
