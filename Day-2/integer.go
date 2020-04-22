package main

import ("fmt"
        "reflect"
        "unsafe")

func main(){
var(
a int8
b int16
c int32
d int64
e int
f uint8
g uint16
h uint32
j uint64
)

fmt.Printf("type of a is %s and Size is ", reflect.TypeOf(a))
fmt.Printf("%d bytes\n", unsafe.Sizeof(a))

fmt.Printf("type of b is %s and Size is ", reflect.TypeOf(b))
fmt.Printf("%d bytes\n", unsafe.Sizeof(b))

fmt.Printf("type of c is %s and Size is ", reflect.TypeOf(c))
fmt.Printf("%d bytes\n", unsafe.Sizeof(c))

fmt.Printf("type of d is %s and Size is ", reflect.TypeOf(d))
fmt.Printf("%d bytes\n", unsafe.Sizeof(d))

fmt.Printf("type of e is %s and Size is ", reflect.TypeOf(e))
fmt.Printf("%d bytes\n", unsafe.Sizeof(e))

fmt.Printf("type of f is %s and Size is ", reflect.TypeOf(f))
fmt.Printf("%d bytes\n", unsafe.Sizeof(f))

fmt.Printf("type of g is %s and Size is ", reflect.TypeOf(g))
fmt.Printf("%d bytes\n", unsafe.Sizeof(g))

fmt.Printf("type of h is %s and Size is ", reflect.TypeOf(h))
fmt.Printf("%d bytes\n", unsafe.Sizeof(h))

fmt.Printf("type of j is %s and Size is ", reflect.TypeOf(j))
fmt.Printf("%d bytes\n", unsafe.Sizeof(j))

}
