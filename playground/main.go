package main

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) String() string {
	linkedListString := ""
	node := l.head
	for {
		linkedListString += strconv.Itoa(node.data)
		if node.next == nil {
			break
		}
		linkedListString += "->"
		node = node.next
	}
	return linkedListString
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) append(n *node) {
	lastNode := l.head
	for {
		if lastNode.next == nil {
			break
		}
		lastNode = lastNode.next
	}
	lastNode.next = n
}

func main() {
	myList := linkedList{}
	node1 := &node{
		data: 44,
	}
	node2 := &node{
		data: 43,
	}
	node3 := &node{
		data: 41,
	}
	node4 := &node{
		data: 444,
	}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.append(node4)
	fmt.Println(myList.String())

}
