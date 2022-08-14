package main

import "fmt"

type TrieNode struct {
	child [26]*TrieNode
	flag  [26]bool
}

func NewNode() *TrieNode {
	newNode := TrieNode{}
	for i := 0; i < 26; i++ {
		newNode.child[i] = nil
		newNode.flag[i] = false
	}
	return &newNode
}

func NewTrie() *TrieNode {
	root := NewNode()
	return root
}

func Contains(str string, root *TrieNode) bool {
	node := root
	for i := 0; i < len(str); i++ {
		switch i {
		case len(str) - 1:
			if node.flag[str[i]-97] == true && node.child[str[i]-97] != nil {
				return true
			} else {
				return false
			}
		default:
			if node.child[str[i]-97] == nil {
				return false
			} else {
				node = node.child[str[i]-97]
			}
		}
	}
	return false
}

func insert(str string, root *TrieNode) {
	node := root
	for i := 0; i < len(str); i++ {
		switch i {
		case len(str) - 1:
			if node.child[str[i]-97] == nil {
				node.child[str[i]-97] = NewNode()
				node.flag[str[i]-97] = true
			} else {
				node.flag[str[i]-97] = true
			}
		default:
			if node.child[str[i]-97] == nil {
				node.child[str[i]-97] = NewNode()
			}
			node = node.child[str[i]-97]
		}
	}
}

// assumption : the word contains a_z letters. letters must be lower case
func main() {
	trie := NewTrie()
	insert("mohammad", trie)
	fmt.Println(Contains("mohammad", trie))
	insert("moha", trie)
	fmt.Println(Contains("moha", trie))
	fmt.Println(Contains("mohamm", trie))
}
