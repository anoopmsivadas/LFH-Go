package main

import ("fmt"
        "reflect"
        "unsafe")

func main(){
var(
a float32 = 3
b float32 = 5
c=complex(a,b)
d complex64 = 4 + 5i
e complex128
)

fmt.Printf("type of c is %s and Size is ", reflect.TypeOf(c))
fmt.Printf("%d bytes\n", unsafe.Sizeof(c))

fmt.Printf("type of d is %s and Size is ", reflect.TypeOf(d))
fmt.Printf("%d bytes\n", unsafe.Sizeof(d))

fmt.Printf("type of e is %s and Size is ", reflect.TypeOf(e))
fmt.Printf("%d bytes\n", unsafe.Sizeof(e))

fmt.Println("c:",c)
fmt.Println("d:",d)
fmt.Println("e:",e)

fmt.Println(c+d, c-d, c*d, c/d)

}
