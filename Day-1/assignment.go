package main

import (
	"fmt"
)

func main() {
        var(
            s int
            t int)
	fmt.Println("Enter first num")
	fmt.Scanln(&s)
	fmt.Println("Enter second num")
	fmt.Scanln(&t)
	fmt.Println("Sum:",s+t)
}

