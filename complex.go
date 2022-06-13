//go 1.10.4

package main

import (
  "fmt"
  "reflect"
 // "unsafe"
)


func main() {
  
  var r float32=2
  var i float32=5
  
  c1:=complex(r, i)  //Initialisation method 1
  
  //Initialisation method 2
  var c2 complex64
  c2 = 6 - 9i
  
  //Display size
  //fmt.Printf("c1's size is %d bytes\n",unsafe.SizeOf(c1))
  //fmt.Printf("c2's size is %d bytes\n",unsafe.SizeOf(c2))

  //Display Type
  fmt.Printf("c1's type is %s\n",reflect.TypeOf(c1))
  fmt.Printf("c2's type is %s\n",reflect.TypeOf(c2))


  //Complex numbers arithmetic
  fmt.Printf("Addition %f",c1+c2)
}