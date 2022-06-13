//go 1.10.4

package main

import "fmt"

func main() {

  
  //s := make([]int,4,2)
  //Error length > capacity
  
  s1 := make([]int,1,2)
  
  fmt.Println(s1)  // [0]
  
  s1 = append(s1,2)
  fmt.Println("Slice is ",s1)   //[0 2]
  fmt.Println("Lenth ",len(s1))    //2
  fmt.Println("Capacity ",cap(s1)) //2
  
  
   s1 = append(s1,3)
  fmt.Println("Slice is ",s1)   //[0 2 3]
  fmt.Println("Lenth ",len(s1))    //3
  fmt.Println("Capacity ",cap(s1)) //4
}