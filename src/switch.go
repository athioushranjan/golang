package main
import "fmt"
func main(){
  // declaring the variable
  var input int
  fmt.Println("enter the value between 0-9")
  fmt.Scanf("%d", &input)
  if input < 10 {
    switch input % 2  {
    case 0 :
        fmt.Println("number is even")
    default :
      fmt.Println("number is odd")
    }

  }else{
    fmt.Println("enter valid number between 0-9")
  }

}
