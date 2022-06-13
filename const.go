//go 1.10.4

package main

import (
  "fmt"
  "time"
)

func main() {
  
  const ( 
    e =2.71828182845904523536028747135266249775724709369995957496696763
    pi = 3.14159265358979323846264338327950288419716939937510582097494459 
  ) //sequence of constants in one declaration
  
  fmt.Println("e =",e,"\n","pi =",pi)
  
  //const id TYPE
  const noDelay time.Duration = 0 //time.Duration is a named type whose underlying type is int64
  const timeout = 5 * time.Minute //time.Minute is a constant of time.Duration
  
  
  fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
  fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s 
  fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

  const pi2 = 2 * pi
  fmt.Printf("%T %[1]v\n", pi2) 
  
  const ( 
    a=1 
    b 
    c
    d 
  ) // right-handside expressionmay be omitted for all 
    // but the first of the group, implying
    // that the previous expression and its type 
    // should be used again
  fmt.Println(a, b, c, d) // "1 1 2 2"
    fmt.Printf("%T %[1]v\n", a)
    fmt.Printf("%T %[1]v\n", d)

}