package main

import "fmt"

func main(){
  
  //Creating a slice
  slice := make([]int,5)
  fmt.Println("Slice is",slice)
  
  slice[0] = 1
  slice[1] = 2
  
  fmt.Println("Slice[0] is",slice[0])
  fmt.Println("Slice[1] is",slice[1])
  
  fmt.Println("and Slice is",slice)
  
  //Length of slice
  fmt.Println("Length of slice is",len(slice))
  
  //Capacity of slice
  fmt.Println("Capacity of slice is",cap(slice))
  
  //using append
  slice = append(slice,6)
  fmt.Println("Now Slice is",slice)
  fmt.Println("and Length of slice is",len(slice))
  
  slice = append(slice,7,8,9)
  fmt.Println("Now Slice is",slice)
  fmt.Println("and Length of slice is",len(slice))
  fmt.Println("and Capacity of slice is",cap(slice))
  
  //Copying slice
  slice2 := make([]int,len(slice))
  copy(slice2,slice)
  fmt.Println("New Slice is",slice2)
  
}