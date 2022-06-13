package main
import (
  "fmt"
  "os"
)


func add_comp(c1, c2 complex128) complex128  {
  return c1 + c2
}

func sub_comp(c1, c2 complex128) complex128 {
  return c1 - c2
}

func bigger_one(c1, c2 complex128) complex128 {
  if (real(c1) <= real(c2) && imag(c1) < imag(c2)) {
    return c2
  } else if c1 == c2 {
    return c2
  } else {
    return c1
  }
}



func main()  {
  fmt.Println("For 1st number")
  fmt.Println("Enter real part")
  var real_1 float64
  fmt.Scan(&real_1)
  fmt.Println("Enter the imaginary part")
  var imag_1 float64
  fmt.Scan(&imag_1)

  fmt.Println("Enter the 2nd number")
  fmt.Println("Enter real part")
  var real_2 float64
  fmt.Scan(&real_2)
  fmt.Println("Enter the imaginary part")
  var imag_2 float64
  fmt.Scan(&imag_2)

  c1 := complex(real_1, imag_1)
  c2 := complex(real_2, imag_2)

  // var choice string
  exit_choice := 44
  // var choice int

  // reader := bufio.NewReader(os.Stdin)

  for exit_choice != 4{
    var choice int
       fmt.Println("Your Calculator")
       fmt.Println("Press '1' for addition")
       fmt.Println("Press '2' for substraction")
       fmt.Println("Press '3' for bigOf")
       fmt.Println("Press '4' for exit")
       fmt.Println("Enter your choice :")

       fmt.Scanln(&choice)

       if choice == 1{
         fmt.Println(add_comp(c1, c2))
       } else if choice == 2 {
          fmt.Println(sub_comp(c1, c2))
       } else if choice == 3 {
          fmt.Println(bigger_one(c1, c2))
       } else if choice == 4{
         fmt.Println("Exiting...")
         // exit_choice = 4
         os.Exit(0)
       } else {
         fmt.Println("Wrong CHoice")
       }

    // switch choice {
    // case 1:
    //      fmt.Println(add_comp(c1, c2))
    // case 2:
    //       fmt.Println(sub_comp(c1, c2))
    // case 3:
    //       fmt.Println(bigger_one(c1, c2))
    // case 4:
    //       fmt.Println("Exiting...")
    //       os.Exit(0)
    // default:
    //     fmt.Printf("Wrong choice...")
    //   }
    }
}
