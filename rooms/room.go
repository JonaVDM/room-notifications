package rooms

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const URL string = "https://roomapi.hexia.io/api/v1/actueel-aanbod?limit=30&locale=nl_NL&page=0&sort=%2BreactionData.aangepasteTotaleHuurprijs"

func GetRooms() []Datum {
	body := []byte(`{
    "filters": {
      "regio.id": {
        "$eq": 8
      }
    } 
  }`)

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	var r Thing
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return r.Data
}
