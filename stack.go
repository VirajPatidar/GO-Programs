package main

import (
	"fmt"
)

type Stack struct {
	Elements []int
}

func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

func (s *Stack) Pop() (int) {
	if s.IsEmpty() {
		return 0
	} else {
		lastElemIndex := len(s.Elements) - 1
		var lastElement int
		lastElement, s.Elements = s.Elements[lastElemIndex], s.Elements[:lastElemIndex]
		return lastElement
	}
}

func (s *Stack) Peek() (int) {
	if s.IsEmpty() {
		return 0
	} else {
		return s.Elements[len(s.Elements)-1]
	}
}

func (s *Stack) Length() int {
	return len(s.Elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Println(stack.IsEmpty())
	
	peek1:= stack.Peek()
	fmt.Println(peek1)
	
	fmt.Println(stack.Length())
	fmt.Println()
	
	elem1 := stack.Pop()
	fmt.Println(elem1)
	
	elem2 := stack.Pop()
	fmt.Println(elem2)
	
	elem3 := stack.Pop()
	fmt.Println(elem3)
	
	fmt.Println(stack.IsEmpty())
}