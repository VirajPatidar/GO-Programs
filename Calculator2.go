package main

import (
    "fmt"
    "bufio"
    "os"
)


//func addInt(a int,b int){
func addInt(a,b int) int{
  c := a+b
  return c
  
}

func subInt(a,b int) int{
  c := a-b
  return c
  
}

func multInt(a,b int) int{
  c := a*b
  return c
  
}

func modInt(a,b int) int{
  c := a%b
  return c
  
}

type calc func(int, int) int

//Function as parameter
func operate(a int, b int, f calc) int{
  result := f(a,b)
  return result
  
} 

func main() {
    
    var choice rune
    var a,b int
    
    reader := bufio.NewReader(os.Stdin)
       
    for choice != 'e'{
         fmt.Println("Your Calculator")
         fmt.Println("Press '+' for addition")
         fmt.Println("Press '-' for substraction")
         fmt.Println("Press '*' for multiplication")
         fmt.Println("Press '%' for Modulus")
         fmt.Println("Press 'e' for exit")
         fmt.Println("Enter your choice :")
        
         choice,_ ,_= reader.ReadRune()
        
        switch choice {
      case '+':
           fmt.Println("Enter 2 values to add")
           fmt.Scanf("%d %d",&a,&b)
           fmt.Println("Addition :",operate(a,b, addInt))
      case '-':
            fmt.Println("Enter 2 values to sub")
            fmt.Scanf("%d %d ",&a,&b)
        fmt.Println("Substraction :",operate(a,b, subInt))
        case '*':
            fmt.Println("Enter 2 values to mult")
            fmt.Scanf("%d %d ",&a,&b)
        fmt.Println("Multiplication :",operate(a,b, multInt))
      case '%':
            fmt.Println("Enter 2 values to mod")
            fmt.Scanf("%d %d",&a,&b)
        fmt.Println("Modulus :",operate(a,b, modInt))
        case 'e':
            fmt.Println("Exiting...")
            os.Exit(0)
        default:
        fmt.Printf("Wrong choice...")
      }  
    }
 
}