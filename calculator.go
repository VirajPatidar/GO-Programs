package main

import "fmt"

func main() {
	var value int

    fmt.Println("Enter your choice: ")
	fmt.Scanln(&value)

	p := 35
	q := 20

	switch value{
	case 1:
		result1 := p + q
		fmt.Printf("Result of p + q = %d\n", result1)
	case 2:
		result2 := p - q
		fmt.Printf("Result of p - q = %d\n", result2)
	case 3:
		result3 := p * q
		fmt.Printf("Result of p * q = %d\n", result3)
	case 4:
		result4 := p / q
		fmt.Printf("Result of p / q = %d\n", result4)
	case 5:
		result5 := p % q
		fmt.Printf("Result of p %% q = %d\n", result5)
	default:
		fmt.Println("Invalid")
	}

}
