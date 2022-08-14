package main

import (
	"fmt"
	"math"
)

type BNode struct {
	leftChild  *BNode
	rightChild *BNode
	data       int
}

func NewBinaryTree() *BNode {
	root := BNode{
		leftChild:  nil,
		rightChild: nil,
		data:       math.MaxInt32,
	}
	return &root
}

func NewBNode(value int) *BNode {
	node := BNode{
		leftChild:  nil,
		rightChild: nil,
		data:       value,
	}
	return &node
}

func insert(value int, node *BNode) {
	if node != nil {
		if node.data == math.MaxInt32 {
			node.data = value
		} else if value < node.data {
			if node.leftChild == nil {
				node.leftChild = NewBNode(value)
			} else {
				insert(value, node.leftChild)
			}
		} else {
			if node.leftChild == nil {
				node.rightChild = NewBNode(value)
			} else {
				insert(value, node.rightChild)
			}
		}
	} else {
		node = NewBNode(value)
	}
}

func search(value int, node *BNode) bool {
	if node == nil {
		return false
	}
	if node.data == value {
		return true
	}
	if value > node.data {
		return search(value, node.rightChild)
	} else if value < node.data {
		return search(value, node.leftChild)
	}
	return false
}

func PreOrderTraversal(node *BNode) {
	if node != nil {
		fmt.Println(node.data)
		PreOrderTraversal(node.leftChild)
		PreOrderTraversal(node.rightChild)
	}
}

func main() {
	BT := NewBinaryTree()
	insert(32, BT)
	insert(44, BT)
	insert(22, BT)
	insert(23, BT)
	insert(45, BT)
	insert(21, BT)
	insert(57, BT)
	fmt.Println(search(32, BT))
	fmt.Println(search(22, BT))
	fmt.Println(search(44, BT))
	fmt.Println(search(56, BT))
	PreOrderTraversal(BT)
}
