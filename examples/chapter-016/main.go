package main

import "fmt"

func main() {
	// short decoration operator is only for function
	y := 42
	z := 42.0

	fmt.Printf("%v of type %T\n", y, y)
	fmt.Printf("%v of type %T\n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T\n", m, m)

	// this is not work - in go you cannot take a VALUE that is float32
	// and store it in a variable that is declared to hold a VALUE of float64
	// z = m

	// this is work
	z = float64(m)
	fmt.Printf("%v of type %T\n", z, z)
}
