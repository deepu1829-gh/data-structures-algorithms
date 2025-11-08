package main

import "fmt"

// Node structure
type Node struct {
	value int
	left  *Node
	right *Node
}

func inorder(root *Node) []int {
	result := []int{}
	fillInorder(root, &result)
	return result
}

// Helper recursive function
func fillInorder(root *Node, result *[]int) {
	if root == nil {
		return
	}
	fillInorder(root.left, result)        // left
	*result = append(*result, root.value) // visit root
	fillInorder(root.right, result)       // right
}

func newNode(value int) *Node {
	return &Node{value: value}
}

func main() {
	/*
	        4
	      /   \
	     2     6
	    / \   / \
	   1  3  5  7
	*/

	root := newNode(4)
	root.left = newNode(2)
	root.right = newNode(6)
	root.left.left = newNode(1)
	root.left.right = newNode(3)
	root.right.left = newNode(5)
	root.right.right = newNode(7)

	result := inorder(root)
	fmt.Println("In-order result:", result)
}
