package main

import (
	"errors"
	"fmt"
)

type DoubleLinkNode struct {
	next *DoubleLinkNode
	prev *DoubleLinkNode
	data int
}

type DoubleLinkList struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
}

func (node *DoubleLinkList) pushFront(new_date int) *DoubleLinkNode {
	ptr := &DoubleLinkNode{data: new_date}
	ptr.next = node.head
	if node.head != nil {
		node.head.prev = ptr
	} else {
		node.tail = ptr
	}
	node.head = ptr

	return ptr
}

func (node *DoubleLinkList) pushBack(new_date int) *DoubleLinkNode {
	ptr := &DoubleLinkNode{data: new_date}
	if node.tail != nil {
		node.tail.next = ptr
		ptr.prev = node.tail
	} else {
		node.head = ptr
	}
	node.tail = ptr

	return ptr
}

func (node *DoubleLinkList) popFront() {
	ptr := node.head.next
	if ptr != nil {
		ptr.prev = nil
	} else {
		node.tail = nil
	}

	node.head = ptr
}

func (node *DoubleLinkList) popBack() {
	ptr := node.tail.prev
	if ptr != nil {
		ptr.next = nil
	} else {
		node.head = nil
	}

	node.tail = ptr
}

func (node *DoubleLinkList) getElemById(index int) (int, error) {
	ptr := node.head
	for n := 0; n != index; n++ {
		if ptr == nil {
			return 0, errors.New("в массиве нет данного индекса")
		}
		ptr = ptr.next
	}

	return ptr.data, nil
}

func (node *DoubleLinkList) print() {
	for ptr := node.head; ptr != nil; ptr = ptr.next {
		fmt.Println(ptr.data)
	}
}

func (node *DoubleLinkList) printReversed() {
	for ptr := node.tail; ptr != nil; ptr = ptr.prev {
		fmt.Println(ptr.data)
	}
}

func main() {
	lst := DoubleLinkList{}
	lst.pushBack(1)
	lst.pushBack(3)
	lst.pushBack(4)
	lst.pushBack(5)
	lst.pushBack(2)

	lst.print()

	lst.popFront()

	fmt.Println("______________________")

	lst.print()

	elem, err := lst.getElemById(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Значение элемента:", elem)
	}

	lst.printReversed()
}
