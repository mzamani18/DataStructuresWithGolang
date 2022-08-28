# important data structures with Golang

in this repo i implement Trie and LinkList and BinaryTree data structures.

#### Trie
```
1. NewNode : for creating a new node in trie or make a root for trie.
2. insert : for inserting a new Word to Trie. 
3. contains : for checking a word exists or not.

```

#### LinkList
```
1. NewLinkList : for creating a new link list.
2. insertNode : for inserting a data to link list. it insert at back.
3. GeneralDelete : for deleting a unknown position data.
4. DeleteFirst : deleting first node. implemented in O(1).
5. DeleteLast : deleting last node. impelemented in O(n).
6. print : for printing link list.
```

#### BinaryTree
```
1. NewBinaryTree : for creating new BinaryTree
2. insert : for inserting a new data to our binary tree.
3. search : search a data in O(log(n)) and this function return a bool.
4. PreOrderTraversal : print data of tree in pre order traversal.
```


#### Stack 
```
1. NewStack : for creating new Stack
2. Push : push data to last of the Stack in o(1)
3. Pop : delete and return last element of Stack in O(1)
4. Top : return last element that exist in the Stack . if the stack be empty it will return -1.
5. Len : return size of Stack.
```