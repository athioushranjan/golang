package main
import "fmt"
func main(){
  // variable declared will be used to get the input from user
  var age int
  //getting input from the user
fmt.Scanf("%d", &age)
//checking the input value according to logic for voting
if age >= 18 {
  fmt.Println("your age is",age, "\nyou are capable for voting")
}else{
  fmt.Println("your age is",age,"\nyou are not capable for voting")
}
}
