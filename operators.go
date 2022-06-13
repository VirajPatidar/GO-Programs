package main

import "fmt"

func main() {
  var x, y = 4, 5

    // Arithmetic Operators
  fmt.Printf("x + y = %d\n", x+y)   // Addition
  fmt.Printf("x - y = %d\n", x-y)   // Substraction
  fmt.Printf("x * y = %d\n", x*y)   // Multiplication
  fmt.Printf("x / y = %d\n", x/y)   // Division
  fmt.Printf("x mod y = %d\n", x%y) // Modulus

  x++                               // Increment 
  fmt.Printf("x++ = %d\n", x)           

  //fmt.Printf("y-- = %d\n", y--)   ERROR
    y--
    fmt.Printf("y-- = %d\n", y)
    
    // Multiply and asign
    x *= y
  fmt.Println("*=", x)
    
    // Comparison / Relational operators
    // returns TRUE or FALSE depending in result of comparison
    fmt.Println(x == y)
  fmt.Println(x != y)
   
    // Logical Operators
    // returns TRUE or FALSE depending in result of operation
    var z = 6
    fmt.Println(x < y && x > z)
  fmt.Println(x < y || x > z)
  fmt.Println(!(x == y && x > z))
    
    z = x & y
  fmt.Println("x & y  =", x & y)

  z = x | y
  fmt.Println("x | y  =", z)

  z = x ^ y
  fmt.Println("x ^ y  =", z)

  z = x << 1
  fmt.Println("x << 1 =", z)

  z = x >> 1
  fmt.Println("x >> 1 =", z)
    
    //var x1, z1 uint=4,3
    //z1 = x1 & y        ERROR Types not same
    // fmt.Println("x1 & y  =", z1)
}