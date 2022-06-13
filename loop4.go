package main

import "fmt"

func main() {
  
  m := map[int]string{
    1 : "one",
    2 : "two",
    3 : "three",
    4 : "four",
    5 : "five",
  }
  
  for key, val := range m {
    fmt.Println(key, val)
  } 


}