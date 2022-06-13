//go 1.10.4

package main

import (
  "fmt"
  "reflect"
)

func main() {
  var b byte='m'
  r :='₹'
  
  //Printing types
  fmt.Printf("b's type os %s\n",reflect.TypeOf(b))
  fmt.Printf("r's type os %s\n",reflect.TypeOf(r))
  fmt.Println(reflect.TypeOf(r))
  
  //Printing values
  fmt.Printf("b's value is %c\n",b)
  fmt.Printf("r's value is %c\n",r)
  
  //Printing Unicode Code Point
  fmt.Printf("Unicode Code Point : %U",r)
}