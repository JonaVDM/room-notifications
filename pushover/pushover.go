package pushover

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

const URL string = "https://api.pushover.net/1/messages.json"

func SendNotification(title, message, url string) {
	token := os.Getenv("PUSH_TOKEN")
	user := os.Getenv("PUSH_USER")

	body := []byte(fmt.Sprintf(`{
    "token": "%v",
    "user": "%v",
    "title": "%v",
    "message": "%v",
    "url": "%v"
  }`, token, user, title, message, url))

	_, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	// txt, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// log.Println(string(txt))
}
