package main

import (
	"fmt"
)


type pelli int
var x pelli

func main() {
	fmt.Printf("x: %v type: %T \n", x, x)

	x = 42

	fmt.Printf("x: %v type: %T \n", x, x)
}
