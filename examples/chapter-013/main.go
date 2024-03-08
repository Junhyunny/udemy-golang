package main

import "fmt"

func main() {
	// := short decoration operator
	// quickly declare, assign the value to variable
	a := 42
	fmt.Println(a)

	// _ blank identifier
	// We can use the blank identifier to declare and use the unused variables.
	// The unused variables are the variables that are only defined once in the program and not used.
	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	// cannot declare not using variable
	// compile error when not used variables existed to keep code pollution
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// when you want zero value then var keyword, and type uses
	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// then assign a VALUE of that TYPE to the variable
	// initialzing a variable
	var h int = 44
	fmt.Println(h)
}
