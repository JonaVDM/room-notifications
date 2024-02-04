package main

import (
	"fmt"

	"github.com/jonavdm/room-notifications/rooms"
)

const ROOM_URL string = ""

func main() {
	r := rooms.GetRooms()
	fmt.Println(len(r))
}
