package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	title string
	body  string
}

var database []Item

type API int

func (a *API) GetByName(title string, response *Item) error {
	var resultItem Item

	for _, item := range database {
		if title == item.title {
			resultItem = item
			break
		}

	}
	response = &resultItem
	return nil
}

func (a *API) AddItem(item Item, response *Item) error {
	database = append(database, item)
	response = &item
	return nil
}

func (a *API) EditItem(edit Item, response *Item) error {
	for index, item := range database {
		if edit.title == item.title {
			database[index] = edit
			break
		}
	}
	response = &edit
	return nil
}

func DeleteItem(delItem Item, response *Item) error {
	for index, item := range database {
		if item.title == delItem.title && item.body == delItem.body {
			database = append(database[:index], database[index+1:]...)
			break
		}
	}
	response = &delItem
	return nil
}

func main() {
	var api = new(API)
	if err := rpc.Register(api); err != nil {
		log.Fatal("Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("Serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving rpc: ", err)
	}
}
