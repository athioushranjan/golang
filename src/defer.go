package main
import "fmt"
func main(){
  for i := 0 ; i < 6 ; i++ {
    //deferred calls will be executed in last-in-first-out order
    defer fmt.Println("value of i ", i )
  }
  fmt.Println("go with deferred")
  
}
