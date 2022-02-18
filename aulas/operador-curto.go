package main

import (
	"fmt"
)

var global = "global"
// global := "global" essa declaração não funciona no escopo global

func main() {


	// :=(marmota) operador de declaração
	x:= 10
	y:= "string"
	z:= 10.0

	fmt.Printf("x: %v, %T \n", x,x)
	fmt.Printf("y: %v, %T \n", y,y)
	fmt.Printf("z: %v, %T \n", z,z)

	// = operador de atribuição
	x = 15
	fmt.Printf("x: %v, %T \n", x,x)

	// O operador de declaração deve declara apenas 1 variavel e pode atribuir outras

	x,m := 20, 10
	fmt.Printf("x: %v, %T \n", x,x)
	fmt.Printf("m: %v, %T \n", m,m)

	fmt.Printf("global: %v, %T \n", global, global)
}
