package main
import "fmt"
func main(){
  array := []int{5,8,6,7,9,1,2}
  //doing iteration on the array
  for _, value := range array {
    fmt.Println("values", value)
  }
  // i will be use to get the current index of the array
  for i, value := range array {
    fmt.Println(value, i)
  }
}
