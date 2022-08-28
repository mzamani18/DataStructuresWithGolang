package main

import "fmt"

type Stack struct {
	top    *node
	length int
}

type node struct {
	data int
	prv  *node
}

func NewStack() *Stack {
	return &Stack{
		nil,
		0,
	}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Top() int {
	if s.length == 0 {
		return -1
	}
	return s.top.data
}

func (s *Stack) Push(value int) {
	tmp := s.top
	s.top = &node{
		data: value,
		prv:  tmp,
	}
	s.length++
}

func (s *Stack) Pop() int {
	if s.length == 0 {
		return -1
	}
	tmp := s.top
	s.top = s.top.prv
	s.length--
	return tmp.data
}

func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	fmt.Println(s)
	fmt.Println(s.length)
	s.Pop()
	fmt.Println(s)
	fmt.Println(s.length)
	fmt.Println(s.Top())
}
