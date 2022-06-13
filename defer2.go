//go 1.10.4

package main
import "fmt"

func greetUser(s string){
  fmt.Println("Good",s)
}

func main() {
  fmt.Println("Its morning...")
  
  defer greetUser("morning")
  
  fmt.Println("Its afternoon...")
  
  defer greetUser("afternoon")
  
  fmt.Println("Its evening...")
  
  defer greetUser("evening")
  
  fmt.Println("Thats defer in Golang...using stack")
}