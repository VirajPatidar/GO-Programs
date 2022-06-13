package main
 
import (
    "fmt"
)
 
func t(i interface{}) string {
    switch i.(type) {
        case string:
        return "A string value"
        case int:
        return "A number"
        default:
        return "Other"
    }
}
 
func num(){
  var a interface{} = 5
  fmt.Println(t(a))// a number
}

func ot(){
  var a interface{} = true
  fmt.Println(t(a))  //a boolean
}


func main() {
    var a interface{} = "A sample string"
    fmt.Println(t(a)) // A string value
    num()
    ot()
}