package main

import "fmt"

type point struct {
	x, y int
}

type nums struct {
	a, b int
}

func (pt point) display() string {
	s3 := fmt.Sprintf("%v, %v", pt.x, pt.y)
	return s3
}

func (n nums) display() string {
	s1 := string(n.a)
	s2 := string(n.b)
	s3 := fmt.Sprintf("%v, %v", s1, s2)
	return s3
}


func main() {
	pt := point{
		x: 10,
		y: 7,
	}
	n := nums{
		a: 67,
		b: 72,
	}
	fmt.Println(pt.display())
	fmt.Println(n.display())
}
