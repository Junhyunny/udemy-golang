package main

import "fmt"

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b\n", adams)
	fmt.Printf("42 as heximal, %x\n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f, g, h := 0, 1, 2, 3, 4, 5, 15, 16

	fmt.Printf("%v \t %b \t %#X \n", a, a, a)
	fmt.Printf("%v \t %b \t %#X \n", b, b, b)
	fmt.Printf("%v \t %b \t %#X \n", c, c, c)
	fmt.Printf("%v \t %b \t %#X \n", d, d, d)
	fmt.Printf("%v \t %b \t %#X \n", e, e, e)
	fmt.Printf("%v \t %b \t %#X \n", f, f, f)
	fmt.Printf("%v \t %b \t %#X \n", h, h, h)
	fmt.Printf("%v \t %b \t %#X \n", g, g, g)
}