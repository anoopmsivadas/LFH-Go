package main

import ("fmt"
        "reflect"
        "unsafe")

func main(){
var(
a float32 = 3
b float64 = 5
c = 2.3
)

fmt.Printf("type of a is %s and Size is ", reflect.TypeOf(a))
fmt.Printf("%d bytes\n", unsafe.Sizeof(a))

fmt.Printf("type of b is %s and Size is ", reflect.TypeOf(b))
fmt.Printf("%d bytes\n", unsafe.Sizeof(b))

fmt.Printf("type of c is %s and Size is ", reflect.TypeOf(c))
fmt.Printf("%d bytes\n", unsafe.Sizeof(c))

}
