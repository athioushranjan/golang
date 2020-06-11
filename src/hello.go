package main
import "fmt"
func main() {
  var value int = 5 ;
  var name string = "Lets play with go" ;
  var decimal float32 = 5.678614 ;
    fmt.Printf("hello, world\n") ;
    fmt.Printf("%d\n ", value);
  fmt.Println(name) ;
  fmt.Println(decimal) ;
  fmt.Println(true || false)
  var initial = "welcome to atom"
  fmt.Println(initial)
  var a, b int = 1,2
  fmt.Println(a+b)
  var valid = true
  fmt.Println(!valid)
  c := 6
  fmt.Println(c)
  const n = 25
  fmt.Println(n)
  const d = 123 / n
  fmt.Println(float32 (d))
  const s = "string"
  fmt.Println(s)
  i := 2
  j := 3
  for{
    i++
    j++
    if i == 6 && j < 10  {
      fmt.Println("the break value of i =", i)
      break ;
    }else{
      fmt.Println("the value of i = ", i , "\tthe value of j= ", j )
    }
    switch i {
    case 3 , 4 :
      fmt.Println("value is", i)
      default :
      fmt.Println("sum of value",i, j,"=" , i+j )
    }
  }
  var x[5] int
  fmt.Println("emp:", x)
  x[3] = 100
  fmt.Println("set:",x)
  i = 0
  for  i< 5 {
    fmt.Printf("value at position %d = %d \n", i ,x[i])
    i++
  }
  string1 , string2 := swap("hello", "world")
  fmt.Println("swap result", string1, string2)
  num1, num2 := split(16)
  fmt.Println("split function value", num1, num2)


}
func swap(x, y string) (string, string){
  return y, x
}
func split(sum int ) (x, y int ){
  x = sum - 10
  y = sum - x
  return
}
