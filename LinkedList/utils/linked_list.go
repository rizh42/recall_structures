package LinkedList

import (
	"fmt"
)

type Node struct {
	Height int
	Width  int
	Next   *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (list *LinkedList) InsertAtHead(Height int, Width int) {
	node := &Node{Height, Width, nil}

	if list.Head == nil {
		list.Head = node
	} else {
		tmp := list.Head
		list.Head = node
		node.Next = tmp
	}

	list.Length++
}

func (list *LinkedList) InsertAtTail(Height int, Width int) {
	node := &Node{Height, Width, nil}

	if list.Head == nil {
		list.Head = node
	} else {
		tmp := list.Head
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = node
	}

	list.Length++
}

func (list *LinkedList) InsertAtPos(Height int, Width int, pos int) {
	if pos == 0 {
		list.InsertAtHead(Height, Width)
	} else if pos == list.Length-1 {
		list.InsertAtTail(Height, Width)
	} else {
		node := &Node{Height, Width, nil}
		tmp := list.Head

		for i := 0; i < pos-1; i++ {
			tmp = tmp.Next
		}

		node.Next = tmp.Next
		tmp.Next = node
		list.Length++
	}
}

func (list *LinkedList) DeleteAtHead() {
	tmp := list.Head.Next
	list.Head = tmp
	list.Length--
}

func (list *LinkedList) DeleteAtTail() {
	tmp := list.Head
	var tmp2 *Node
	for tmp.Next != nil {
		tmp2 = tmp
		tmp = tmp.Next
	}
	tmp2.Next = nil
	list.Length--
}

func (list *LinkedList) DeleteAtPos(pos int) {
	if pos == 0 {
		list.DeleteAtHead()
	} else if pos == list.Length-1 {
		list.DeleteAtTail()
	} else {
		tmp := list.Head

		for i := 0; i < pos-1; i++ {
			tmp = tmp.Next
		}

		tmp2 := tmp.Next
		tmp.Next = tmp2.Next
		list.Length--
	}
}

func (list *LinkedList) Print() {
	for element := list.Head; element != nil; element = element.Next {
		fmt.Println("Height:", element.Height, "Width:", element.Width)
	}
}

func (list *LinkedList) PrintPos(pos int) {
	element := list.Head
	for i := 0; i < pos-1; i++ {
		element = element.Next
	}

	fmt.Println("Height:", element.Height, "Width:", element.Width)

}
