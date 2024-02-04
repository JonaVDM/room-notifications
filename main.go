package main

import (
	"fmt"

	"github.com/jonavdm/room-notifications/list"
	"github.com/jonavdm/room-notifications/pushover"
	"github.com/jonavdm/room-notifications/rooms"
)

const ROOM_URL string = ""

func main() {
	r := rooms.GetRooms()

	for _, room := range r {
		if ok := list.IsInList(room.ID); !ok {
			var amount int64
			if room.AantalMedebewoners == nil {
				amount = 0
			} else {
				amount = *room.AantalMedebewoners
			}
			message := fmt.Sprintf(`%v %v
Size: %v
Housemates: %v
Rent: %v (%v)
`, room.Street, room.HouseNumber+" "+room.HouseNumberAddition, room.AreaDwelling, amount, room.TotalRent, room.NetRent)
			pushover.SendNotification("New room!", message, "https://www.room.nl/aanbod/studentenwoningen/details/"+room.URLKey)
			list.Add(room.ID)
		}
	}
}
