package main
import "fmt"

func recurSum(n int) int {
	if (n <= 1) {
        return n
	}
    return n + recurSum(n - 1)
}

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)
	fmt.Printf("The sum of first %d numbers is: ", n)
	fmt.Println(recurSum(n))
}