package main

import "fmt"

// Item is a (title, body) structure
type Item struct {
	title string
	body  string
}

var dataBase []Item

// GetByName returns an Item based on its name
func GetByName(name string) Item {
	var getItem Item
	for _, val := range dataBase {
		if val.title == name {
			getItem = val
		}
	}
	return getItem
}

// AddItem store an Item in the database
func AddItem(item Item) Item {
	dataBase = append(dataBase, item)
	return item
}

// EditItem puts the new Item with the title key
func EditItem(title string, edit Item) Item {
	var changed Item
	for idx, val := range dataBase {
		if val.title == title {
			dataBase[idx] = edit
			changed = edit
		}
	}
	return changed
}

// DeleteItem deletes an Item from the database
func DeleteItem(item Item) Item {
	var del Item
	for idx, val := range dataBase {
		if val.title == item.title && val.body == item.body {
			dataBase = append(dataBase[:idx], dataBase[idx+1:]...)
			del = item
			break
		}
	}
	return del
}

func main() {
	fmt.Println("initial database : ", dataBase)
	a := Item{"first", "a test item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}
	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("second database : ", dataBase)
	DeleteItem(b)
	fmt.Println("third database : ", dataBase)
	EditItem("third", Item{"fourth", "a new item"})
	fmt.Println("fourth database : ", dataBase)

	x := GetByName("fourth")
	y := GetByName("first")
	fmt.Println(x, y)

}
