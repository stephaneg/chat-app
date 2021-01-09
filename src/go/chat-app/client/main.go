package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Item is a (title, body) structure
type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("connection error :", err)
	}

	a := Item{"First", "a first item"}
	b := Item{"Second", "a second item"}
	c := Item{"Third", "a third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("current database", db)

	client.Call("API.EditItem", Item{"Second", "a new second item"}, &reply)
	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("after changes database", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item : ", reply)

}
