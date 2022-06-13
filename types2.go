//go 1.10.4

package main

import(
  "fmt"
  "math/bits"
  "reflect"
)

func main() {
  var(
     a int=2 
     b float32=2
     s string="my"
  )
  
  bitsSizeOfInt := bits.UintSize
  
  
  fmt.Printf("%d and %f and %s\n",a,b,s)
  fmt.Printf("Int bits are %d\n",bitsSizeOfInt)
  fmt.Printf("a's type is %s\n",reflect.TypeOf(a))
  fmt.Printf("b's type is %s\n",reflect.TypeOf(b))
  fmt.Printf("s's type is %s\n",reflect.TypeOf(s))
  
  d := 5.0 // default float64
 
  fmt.Printf("d's type is %s\n",reflect.TypeOf(d))
  
  
  var decide bool // default false
  
  fmt.Printf("Boolean is %t\n",decide)
  fmt.Printf("decide's type is %s\n",reflect.TypeOf(decide))
}