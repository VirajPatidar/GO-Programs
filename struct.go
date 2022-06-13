//go 1.10.4

package main


import "fmt"

//structure declaration
type student struct{
  name string
  roll int32
  gpa float32
}


func main() {
  
  //Initialisation 1
  student1 := student{"ABC",191401,8.2}
  fmt.Println("Record 1",student1)
  
  //Initialisation 2
  student2 := student{
                name: "PQR",
                roll:191402,
                gpa:7.5}
  fmt.Println("Record 2",student2)
  
  //Initialisation 3
  student3 := student{
                name: "PQR",
                roll:191402,
                }
  fmt.Println("Record 3",student3)
  
  student3.gpa=6.4
  
  fmt.Println("Record 3 GPA is",student3.gpa)
}