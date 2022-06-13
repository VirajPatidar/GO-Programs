//go 1.10.4

package main

import "fmt"

func main() {
  
  var v int = 10
  
  fmt.Printf("\nValue %v",v)
  fmt.Printf("\nValue %d",v)
  fmt.Printf("\nBits %b",v)
  fmt.Printf("\nType %T",v)
  fmt.Printf("\nOctal %o",v)
  fmt.Printf("\nHexadecimal %x",v)
  fmt.Printf("\nSigned value %+d",v)
  fmt.Printf("\nOctal with preceding 0 %#o",v)
  fmt.Printf("\nHex with preceding 0X %#X",v)
  fmt.Printf("\nSpaced Value% d",v)


}