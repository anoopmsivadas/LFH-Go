package main

import ("fmt"
        "reflect"
        "unsafe")

func main(){
var(
r byte = 'r'
b = ":   ;."
)
s:="abc"
fmt.Printf("Size: %d\n", unsafe.Sizeof(r))
fmt.Printf("Size: %d\n", unsafe.Sizeof(b))
fmt.Printf("Type: %s\n", reflect.TypeOf(r))
fmt.Printf("Type: %s\n", reflect.TypeOf(b))
fmt.Printf("Character: %c\n", r)
fmt.Println(byte(r))
fmt.Println([]byte(b))//output as bytes
fmt.Println([]byte(s))// output as bytes
}
