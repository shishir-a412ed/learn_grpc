package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var response Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("rpc connection error: ", err)
	}

	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	client.Call("API.AddItem", a, &response)
	client.Call("API.AddItem", b, &response)
	client.Call("API.AddItem", c, &response)
	client.Call("API.GetDB", "", &db)
	fmt.Println("database: ", db)
}
