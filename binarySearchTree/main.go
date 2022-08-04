package main

import (
	"fmt"
	"math"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Inserting values
func (n *Node) Insert(val int) {
	if n.Value < val {
		// Move Right
		if n.Right == nil {
			n.Right = &Node{Value: val}
		} else {
			n.Right.Insert(val)
		}
	} else if n.Value > val {
		// Move Left
		if n.Left == nil {
			n.Left = &Node{Value: val}
		} else {
			n.Left.Insert(val)
		}
	}
}

//Delete removes the Item with value from the tree
func (n *Node) Delete(value int) {
	n.remove(value)
}

func (n *Node) remove(val int) *Node {

	if n == nil {
		return nil
	}

	if val < n.Value {
		n.Left = n.Left.remove(val)
		return n
	}
	if val > n.Value {
		n.Right = n.Right.remove(val)
		return n
	}

	if n.Left == nil && n.Right == nil {
		n = nil
		return nil
	}

	if n.Left == nil {
		n = n.Right
		return n
	}
	if n.Right == nil {
		n = n.Left
		return n
	}

	smallestValOnRight := n.Right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.Left != nil {
			smallestValOnRight = smallestValOnRight.Left
		} else {
			break
		}
	}

	n.Value = smallestValOnRight.Value
	n.Right = n.Right.remove(n.Value)
	return n
}

// Search Values
// func (n *Node) Search(val int) bool {

// 	if n == nil {
// 		return false
// 	}

// 	if n.Value < val {
// 		//Move right
// 		return n.Right.Search(val)
// 	} else if n.Value > val {
// 		//Move Left
// 		return n.Left.Search(val)
// 	}
// 	return true
// }

func (n *Node) Search(val int) bool {

	if n == nil {
		return false
	}

	switch {
	case val == n.Value:
		return true
	case val < n.Value:
		return n.Left.Search(val)
	default:
		return n.Right.Search(val)
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

//FindMin finds the min element in the given BST
func (n *Node) FindMin() int {
	if n.Left == nil {
		return n.Value
	}
	return n.Left.FindMin()
}

//FindMax finds the max element in the given BST
func (n *Node) FindMax() int {
	if n.Right == nil {
		return n.Value
	}
	return n.Right.FindMax()
}

//Validate BST
func isValidBST(n *Node) bool {
	if n == nil {
		return true
	}

	return checkValidBST(n, math.MinInt32, math.MaxInt32)
}

func checkValidBST(n *Node, min, max int) bool {
	if n == nil {
		return true
	}

	if n.Value <= min || n.Value >= max {
		return false
	}

	return checkValidBST(n.Left, min, n.Value) && checkValidBST(n.Right, n.Value, max)
}

func main() {

	tree := &Node{Value: 100}

	tree.Insert(20)
	tree.Insert(400)
	tree.Insert(20)
	tree.Insert(320)
	tree.Insert(200)
	tree.Insert(90)
	tree.Insert(80)

	if isValidBST(tree) == true {
		fmt.Println("It is a Tree")
	} else {
		fmt.Println("Not a Tree")
	}

	fmt.Println("InOrder")
	tree.printInOrder()
	fmt.Println("\nPreOrder")
	tree.printPreOrder()
	fmt.Println("\nPostOrder")
	tree.printPostOrder()

	fmt.Println(tree.Search(290))

	tree.Delete(320)
	fmt.Println("After Deletion InOrder")
	tree.printInOrder()

	fmt.Println("min is ", tree.FindMin())
	fmt.Println("max is ", tree.FindMax())

}
