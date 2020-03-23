
package main

import ( "fmt"
         "reflect"
       )

func main(){

var n int
var state string
country:="India"
var fname, lname string = "John", "Doe"
p, q, r := 1, 2, 3
n = 10
state = "Kerala"

name:="Ezio"
fmt.Println("I am "+ name +", from " + country)
fmt.Println(n)
fmt.Println(state)
fmt.Println(country)
fmt.Println(fname + lname)
fmt.Println(p + q + r)
fmt.Println(fname, lname)
fmt.Println(reflect.TypeOf(state))
fmt.Println(reflect.TypeOf(n))
}
