package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"

	fmt.Println(i, j)
	fmt.Printf("%T %T\n", i, j)

	fmt.Println(c, python, java)
	fmt.Printf("%T %T %T\n", c, python, java)
}
