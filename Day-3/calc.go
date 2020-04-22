package main

import "fmt"

func main(){
var ch,a,b int
fmt.Println("1. Add\n2. Subtract\n3. Divide\n4. Multiply")
fmt.Println("Enter the choice and press Enter")
fmt.Scanln(&ch)
fmt.Println("Enter the numbers")
fmt.Scanln(&a)
fmt.Scanln(&b)
switch ch{
  case 1:
    fmt.Println("Sum:",a+b)
  case 2:
    fmt.Println("Difference:",a-b)
  case 3:
    fmt.Println("Quotient:",a/b)
  case 4:
    fmt.Println("Product:",a*b)
  default:
    fmt.Println("Invalid")
  }
}
