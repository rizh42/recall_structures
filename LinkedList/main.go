package main

import (
	"fmt"
	LinkedList "utils"
)

func main() {
	var rooms int
	fmt.Scan(&rooms)
	list := LinkedList.LinkedList{nil, 0}

	for i := 0; i < rooms; i++ {
		var height, width int
		fmt.Scan(&height, &width)
		list.InsertAtTail(height, width)
	}
	list.Print()
}
