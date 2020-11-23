package main

import (
	"fmt"
)

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var resultItem Item

	for _, item := range database {
		if title == item.title {
			resultItem = item
			break
		}

	}
	return resultItem
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	for index, item := range database {
		if title == item.title {
			database[index] = edit
			break
		}
	}
	return edit
}

func DeleteItem(delItem Item) Item {
	for index, item := range database {
		if item.title == delItem.title && item.body == delItem.body {
			database = append(database[:index], database[index+1:]...)
			break
		}
	}
	return delItem
}

func main() {
	fmt.Println("Initial database: ", database)

	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("Second database: ", database)

	DeleteItem(b)
	fmt.Println("Third database: ", database)

	EditItem("third", Item{"fourth", "a fourth item"})
	fmt.Println("Fourth database: ", database)

	x := GetByName("fourth")
	y := GetByName("first")
	fmt.Println(x, y)
}
