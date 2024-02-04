package list

import (
	"errors"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const FILE string = "list.yml"

type Item struct {
	Id      int64
	AddedOn int64
}

func IsInList(id int64) bool {
	items := get()

	for _, item := range items {
		if item.Id == id {
			return true
		}
	}

	return false
}

func Add(id int64) {
	items := get()

	item := Item{
		Id:      id,
		AddedOn: time.Now().UTC().Unix(),
	}

	body, err := yaml.Marshal(append(items, item))
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(FILE, body, 0644)
}

func get() (out []Item) {
	file, err := os.ReadFile(FILE)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return []Item{}
	} else if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(file, &out); err != nil {
		log.Fatal(err)
	}

	return out
}
