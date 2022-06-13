//go 1.10.4

package main

import "fmt"

func iPrintHello(s string){
  fmt.Println("Hello,",s)
  
}

//func addInt(a int,b int){
func addInt(a,b int){
  c := a+b
  fmt.Println("Addition =",c)
  
}
func main() {
  iPrintHello("GoLang")
  addInt(2,3)
}