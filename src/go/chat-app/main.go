package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Item is a (title, body) structure
type Item struct {
	Title string
	Body  string
}

// API is a receiver used to elevate functions to methods
type API int

var dataBase []Item

// GetDB returns the full database
func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = dataBase
	return nil
}

// GetByName returns an Item based on its name
func (a *API) GetByName(name string, reply *Item) error {
	var getItem Item
	for _, val := range dataBase {
		if val.Title == name {
			getItem = val
		}
	}
	*reply = getItem

	return nil
}

// AddItem store an Item in the database
func (a *API) AddItem(item Item, reply *Item) error {
	dataBase = append(dataBase, item)
	*reply = item
	return nil
}

// EditItem puts the new Item with the title key
func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, val := range dataBase {
		if val.Title == edit.Title {
			dataBase[idx] = edit
			changed = edit
		}
	}
	*reply = changed
	return nil
}

// DeleteItem deletes an Item from the database
func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range dataBase {
		if val.Title == item.Title && val.Body == item.Body {
			dataBase = append(dataBase[:idx], dataBase[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registring API", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("listener error API", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving:", err)
	}

}
