package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

// insert value
func (n *Node) Insert(value int) {
	//Check Value itself
	if value < n.Value {
		// Insert to left node
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		// Insert to right node
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Inorder ==> Left -- Root -- Right
func (n *Node) printInOrder() {
	// Left
	if n.Left != nil {
		n.Left.printInOrder()
	}
	//Root
	fmt.Println(n.Value)
	// Right
	if n.Right != nil {
		n.Right.printInOrder()
	}
}

// PretOrder ==> Root -- Left -- Right
func (n *Node) printPreOrder() {
	//Root
	fmt.Println(n.Value)
	// Left
	if n.Left != nil {
		n.Left.printPreOrder()
	}
	// Right
	if n.Right != nil {
		n.Right.printPreOrder()
	}
}

// PostOrder ==> Left -- Right -- Root
func (n *Node) printPostOrder() {
	// Left
	if n.Left != nil {
		n.Left.printPostOrder()
	}
	// Right
	if n.Right != nil {
		n.Right.printPostOrder()
	}
	//Root
	fmt.Println(n.Value)
}

// Search for value
func (n *Node) Search(value int) bool {

	if n == nil {
		return false
	}

	switch {
	case value == n.Value:
		return true
	case value < n.Value:
		return n.Left.Search(value)
	default:
		return n.Right.Search(value)
	}
}

func main() {
	tree := &Node{Value: 50}
	tree.Insert(45)
	tree.Insert(40)
	tree.Insert(20)
	tree.Insert(25)
	tree.Insert(35)

	fmt.Println("\nInOrder")
	tree.printInOrder()
	fmt.Println("\nPreOrder")
	tree.printPreOrder()
	fmt.Println("\nPostOrder")
	tree.printPostOrder()

	fmt.Println(tree.Search(15))

}
