package main

import "fmt"

type ListNode struct {
	nextPtr *ListNode
	data    int
}

type LinkList struct {
	firstPtr *ListNode
	lastPtr  *ListNode
	size     int
}

func NewLinkList() *LinkList {
	linkList := LinkList{
		firstPtr: nil,
		lastPtr:  nil,
		size:     0,
	}
	return &linkList
}

func NewListNode(value int) *ListNode {
	ln := ListNode{
		nextPtr: nil,
		data:    value,
	}
	return &ln
}

func insertNode(value int, Linklist *LinkList) {
	if Linklist.firstPtr == nil {
		Linklist.firstPtr = NewListNode(value)
		Linklist.lastPtr = Linklist.firstPtr
		Linklist.size++
	} else {
		newNode := NewListNode(value)
		Linklist.lastPtr.nextPtr = newNode
		Linklist.lastPtr = newNode
		Linklist.size++
	}
}

func GeneralDelete(value int, linkList *LinkList) {
	node := linkList.firstPtr
	Prenode := linkList.firstPtr
	if node.data == value {
		linkList.firstPtr = linkList.firstPtr.nextPtr
		linkList.size--
	} else {
		node = node.nextPtr
		for node != nil {
			if node.data == value {
				if node.nextPtr == nil {
					Prenode.nextPtr = nil
					linkList.lastPtr = Prenode
				} else {
					Prenode.nextPtr = node.nextPtr
				}
				linkList.size--
				break
			}
			Prenode = node
			node = node.nextPtr
		}
	}
}

func DeleteFirst(linkList *LinkList) {
	if linkList.firstPtr == nil {
		fmt.Println("link list is empty")
	} else {
		linkList.firstPtr = linkList.firstPtr.nextPtr
		linkList.size--
	}
}

func DeleteLast(linkList *LinkList) {
	node := linkList.firstPtr
	PreNode := linkList.firstPtr
	if node == nil {
		fmt.Println("link list is empty")
	} else if node.nextPtr == nil {
		linkList.lastPtr = linkList.firstPtr
		linkList.firstPtr.nextPtr = nil
		linkList.size--
	} else {
		node = node.nextPtr
		for node != nil {
			if node == linkList.lastPtr {
				PreNode.nextPtr = nil
				linkList.lastPtr = node
				linkList.size--
				break
			}
			PreNode = node
			node = node.nextPtr
		}
	}
}

func print(linkList *LinkList) {
	node := linkList.firstPtr
	for node != nil {
		fmt.Println(node.data)
		node = node.nextPtr
	}
}

func main() {
	ln := NewLinkList()
	DeleteFirst(ln)
	insertNode(2, ln)
	DeleteFirst(ln)
	insertNode(3, ln)
	insertNode(24, ln)
	insertNode(22, ln)
	insertNode(21, ln)
	insertNode(9, ln)
	print(ln)
	DeleteFirst(ln)
	print(ln)
	DeleteLast(ln)
	print(ln)
	GeneralDelete(22, ln)
	fmt.Println("after last general delete : ")
	print(ln)
}
