//go 1.10.4

package main

import "fmt"


//func addInt(a int,b int){
func addInt(a,b int) int{
  c := a+b
  return c
  
}

func subInt(a,b int) int{
  c := a-b
  return c
  
}

//Function as parameter
func operate(a int, b int, f func(int,int) int) int{
  result := f(a,b)
  return result
  
}

func main() {
  add := addInt(2,3)
  sub := subInt(6,4)
  fmt.Println("Addition =",add)
  fmt.Println("Addition =",sub)
  
  fmt.Printf("Type for addInt is %T",addInt)
}