package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	// tail   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) prinlistData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}
func (l *linkedList) deleteValueWith(value int) {
	previousToDelete := l.head
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	if previousToDelete.next.next == nil {
		return
	}

	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l *linkedList) Reverse() {

	curr := l.head
	var prev *node
	var next *node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func main() {

	mylist := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 40}
	node5 := &node{data: 50}
	node6 := &node{data: 60}
	node7 := &node{data: 70}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)

	mylist.prinlistData()
	mylist.Reverse()
	mylist.prinlistData()
	mylist.deleteValueWith(30)
	mylist.prinlistData()
	mylist.deleteValueWith(10)
	mylist.prinlistData()
	mylist.deleteValueWith(70)
	mylist.prinlistData()

}
