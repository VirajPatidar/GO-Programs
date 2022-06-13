//go 1.10.4

package main
import "fmt"

func sayBye(){
  fmt.Println("Bye Bye")
}

func main() {
  fmt.Println("Hello")
  
  defer sayBye()
  
  fmt.Println("Enjoy Golang...")
}