//go 1.10.4

package main

import "fmt"

//Named return values
func addSubDivInt(a,b int) (add ,sub int,div float32){
  add = a+b
  sub = a-b
  div = float32(a/b)
  
  return //compulsory
  
}
func main() {
  a,s,d := addSubDivInt(2,5)
  fmt.Println("Result is",a,s,d)
}