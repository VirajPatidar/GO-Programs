package main

import (
	"fmt"
)

type Queue struct {
	Elements []int
}

func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() (int) {
	if len(q.Elements) == 0 {
		return 0
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement
}

func (q *Queue) Length() int {
	return len(q.Elements)
}

func (q *Queue) Peek() (int) {
	if q.IsEmpty() {
		return 0
	}
	return q.Elements[0]
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func main() {

	queue := Queue{}

	emp := queue.IsEmpty()
	fmt.Println(emp)
	
	queue.Enqueue(1)
	fmt.Println(queue)
	
	queue.Enqueue(2)
	fmt.Println(queue)
	
	elem := queue.Dequeue()
	fmt.Println(elem)
	
	fmt.Println(queue)
	
	emp = queue.IsEmpty()
	fmt.Println(emp)
	
	l := queue.Length()
	fmt.Println(l)
}