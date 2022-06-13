package main

import "fmt"

//func sum(vals ...int,char byte) (int,byte) {  // .\variadic2.go:5:10: syntax error: cannot use ... with non-final parameter vals
func sum(char byte,vals ...int) (int,byte) { 
	total := 0
	for _, val := range vals {
	total += val
	}
	var byte1 byte
	if total%2 == 0 {
		byte1 = 'e'
	} else {
		byte1 = char
	}
	
	return total, byte1
}

func main(){
	
	//arr:=[5]int{2,3,4,5,6}

	fmt.Println(sum('a',2))
	fmt.Println(sum('b',2,3))
	//fmt.Println(sum(arr))
	//fmt.Println(sum(1,2,5,6,99,a,b,c,d))
}