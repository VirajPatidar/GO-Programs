package main

import "fmt"

func sum(vals ...int) int { // Variadic function, can have multiple values of same type
	total := 0
	for _, val := range vals {
	total += val
	}
	return total
}

func main(){
	
	arr:=[5]int{2,3,4,5,6}

	fmt.Println(sum(2))
	fmt.Println(sum(2,3))
	fmt.Println(sum(int(arr...)))
	fmt.Println(sum(1,2,5,6,99))
}