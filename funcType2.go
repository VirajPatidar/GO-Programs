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

type calc func(int, int) int

//Function as parameter
func operate(a int, b int, f calc) int{
  result := f(a,b)
  return result
  
}

func main() {
  add := operate(2,3, addInt)
  sub := operate(6,4, subInt)
  fmt.Println("Addition =",add)
  fmt.Println("Addition =",sub)
  
  fmt.Printf("Type for addInt is %T",addInt)
  fmt.Printf("Type for operate is %T",operate)
  //fmt.Printf("Type for calc is %T",calc)
}