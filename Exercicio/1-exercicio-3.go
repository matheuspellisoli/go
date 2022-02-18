package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	
	x := 42
	y := "James Bond"
	z := true


	s := fmt.Sprintf("x: %v, y: %v, z: %v\n", x,y,z)

	fmt.Printf(s)
}

//São valores ZERO