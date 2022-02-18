package main

import (
	"fmt"
)

type pelli int

var x pelli
var y int

func main() {
	x = 42
	y = int(x)

	fmt.Printf("x: %v type: %T \n", x, x)
	fmt.Printf("y: %v type: %T \n", y, y)
}

