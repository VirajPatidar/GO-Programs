//go 1.10.4

package main
import "fmt"

func factorial(num int) int {
	if num > 0 {
		return num * factorial(num-1)
	} else {
		return 1 // 0! == 1
	}
}

func main() {
  fact := factorial(5)
  fmt.Println("Factorial is ",fact)
  fact = factorial(0)
  fmt.Println("Factorial is ",fact)
}