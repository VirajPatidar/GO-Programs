package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
  
  sum = 0
  for i := 1; i < 5; i++ { // "i" redefined as loop is scope
    if i%2 != 0 { // skip odd numbers
        continue
    }
    sum += i
  }
  fmt.Println(sum)
  
  i := 0
  for i <= 10 { // for as while
    fmt.Println(i,i*i)
    i++
    
    }
  
}