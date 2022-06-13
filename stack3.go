//go 1.10.4

package main

import "fmt"

type stack struct{
  stk []byte
  top int
}

func (s *stack) Init(){
  s.top=-1
}

func (s *stack) Push(c byte){
  s.top++
  s.stk=append(s.stk,c)
}

func (s *stack) Display(){
  fmt.Println("Stack top is :",s.stk[s.top])
}

func (s *stack) Pop() ([]byte,byte){
  var tmpStk =make([]byte,len(s.stk)-1)
  fmt.Println(s.stk)
  rem := s.stk[s.top]
  
  n := copy(tmpStk,s.stk[0:s.top]) //open last range :(
  s.top--
  fmt.Println(tmpStk,n)
  fmt.Println(rem,"removed, top is",s.top)
  
  return tmpStk,rem
}
func main() {
  var s stack
  var rem byte 
  s.Init()
  fmt.Println(s)
  
  s.Push('a')
  fmt.Println(s)
  
  s.Push('b')
  fmt.Println(s)
  s.Display()
  
  s.stk,rem = s.Pop()
  
  fmt.Println(rem,"removed, top is",s.top)
  
  s.Display()
  fmt.Println(s)
  
  s.Push('b')
  fmt.Println(s)
}