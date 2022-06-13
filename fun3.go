//go 1.10.4

package main

import "fmt"


//func addInt(a int,b int){
func addInt(a,b int) int{
  c := a+b
  return c
  
}
func main() {
  fmt.Println("Addition is",addInt(2,3))
}