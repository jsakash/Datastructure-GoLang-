package main

import "fmt"

type node struct {
	data string
	next *node
}

type singlyLinkedList struct {
	len  int
	head *node
}

func initList() *singlyLinkedList {
	return &singlyLinkedList{}
}

func (s *singlyLinkedList) AddFront(data string) {
	node := &node{
		data: data,
	}

	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.len++
	return
}

func (s *singlyLinkedList) Size() int {
	return s.len
}

func (s *singlyLinkedList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

func main() {
	var myarray [100]string
	var li int
	singleList := initList()
	fmt.Println("Enter the limit of array")
	fmt.Scan(&li)
	for i := 0; i < li; i++ {
		fmt.Scan(&myarray[i])
	}
	fmt.Println("Array elements")
	fmt.Println(myarray)

	for i := 0; i < li; i++ {
		singleList.AddFront(myarray[i])
	}
	fmt.Println("After converting to linkedlist")
	fmt.Printf("Size: %d\n", singleList.Size())

	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()

}
