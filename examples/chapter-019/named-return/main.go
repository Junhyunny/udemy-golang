package main

import "fmt"

// naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))

	var a, b = split(17)
	fmt.Println(a, b)
}
