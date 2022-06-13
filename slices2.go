//go 1.10.4

package main

import "fmt"

func main() {
  
  myArr := [7]byte{'a','b','c','d','e','f','g'}
  fmt.Println(myArr)
  
  //byte to string with slice
  mySlice := string(myArr[:4])
  fmt.Println(mySlice)
}