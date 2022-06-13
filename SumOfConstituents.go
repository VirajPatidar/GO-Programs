package main

import "fmt"

var digiSum int

func digitSum(digiNum int) {
	var digiReminder int

	for digiSum = 0; digiNum > 0; digiNum = digiNum / 10 {
		digiReminder = digiNum % 10
		digiSum = digiSum + digiReminder
	}
	fmt.Println("The Sum of Digits in this Number = ", digiSum)
	fmt.Println()
}

func sumOfcharactersOfString(strData string) {
	sum := 0
	for i := 0; i < len(strData); i++ {
		sum = sum + int(strData[i])
	}
	fmt.Println("The Sum of ASCII values in this String = ", sum)
	fmt.Println()
}


func sliceInput(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return sliceInput(x, err)
}

func sliceSum(array []int) {
	result := 0
	for _, v := range array {
		result += v
	}
	fmt.Println("The Sum of values in this Slice = ", result)
	fmt.Println()

}

func main() {

	var digiNum int
	var str string

	fmt.Print("Enter the Number to find the Sum of Digits: ")
	fmt.Scanln(&digiNum)
	digitSum(digiNum)

	fmt.Print("Enter the String to find Sum of ASCII values: ")
	fmt.Scanln(&str)
	sumOfcharactersOfString(str)

	fmt.Println("Enter the Slice to find the Sum of its values: ")
	x := sliceInput([]int{}, nil)
	sliceSum(x)

}
