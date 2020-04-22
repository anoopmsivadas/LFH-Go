package main

import "fmt"

func main(){
s:=make([]string, 2, 3)
p:=[]string{"a", "b", "c"}
fmt.Println(s)
fmt.Println(p)
p =append(p, "d")
fmt.Println(p)
}
