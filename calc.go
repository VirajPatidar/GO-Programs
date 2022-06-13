package main
 
import "fmt"

func addInt(a,b complex128) complex128{
  c := a+b
  return c
  
}

func subInt(a,b complex128) complex128{
  c := a-b
  return c
  
}

func maxInt(first, second, third, fourth float64) complex128{
  x:= first*first + second*second
  y:= third*third + fourth*fourth
  if x>y {
      return complex(first, second)
  }else {
      return complex(third, fourth)
  }
}
 
func main() {
        fmt.Printf("Enter Real Part of First Number: ")
        var first float64
        fmt.Scanf("%f",&first)    
        fmt.Printf("Enter Imaginary Part of First Number: ")  
        var second float64  
        fmt.Scanf("%f",&second) 
        fmt.Printf("Enter Real Part of Second Number: ")
        var third float64
        fmt.Scanf("%f",&third) 
        fmt.Printf("Enter Imaginary Part of Second Number: ")
        var fourth float64
        fmt.Scanf("%f",&fourth) 
        var fcom complex128 = complex(first, second)
        var scom complex128 = complex(third, fourth)
        fmt.Println("Addition :",addInt(fcom, scom))
        fmt.Println("Subtraction :",subInt(fcom, scom))
        fmt.Println("Bigger Complex number is :",maxInt(first, second, third, fourth))
}
