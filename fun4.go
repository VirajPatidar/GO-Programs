//go 1.10.4

package main

import "fmt"


func addSubDivInt(a,b int) (int,int,int){
  return a+b,a-b,a/b
  
}
func main() {
  add, sub, div := addSubDivInt(2,5)
  fmt.Println("Result is",add,sub,div)
}