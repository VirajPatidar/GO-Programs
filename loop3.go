package main

import "fmt"

func main() {
  
  strings := []string{"here", "we","GO"}
  
  for i  := range strings {
    fmt.Println(i) //Prints 0  1  2 
   
   } 
   
  //  range loop generates two values:
  //  1. the index,
  //  2. data at this index
  
   for _,s  := range strings {
    fmt.Println(s) //Prints string 
   
   } 
   //blank identifier (underscore) can used 
   //for the return value we're not interested in.


}