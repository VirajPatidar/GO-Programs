//go 1.10.4

package main

import "fmt"

func main() {
  myArr := [3]int{1,2,3}
  myArr2 := [...]int{1,2}
  myArr3 := [2]float32{}
  myStrArr := [...]string{"a","b","c","d"}
  
  fmt.Println(myArr)
  fmt.Println(myArr2)
  fmt.Println(myArr3)
  fmt.Println(myStrArr)
  
}
