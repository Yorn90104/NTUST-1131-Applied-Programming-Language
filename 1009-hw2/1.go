package main

import "fmt"

func main() {

	var (
		x  = 3
		y  = 7
		z  = 11
		b1 = x == y
		b2 = x <= z
		b3 = b1 || b2
	)
	fmt.Println(b1, b2, b3)
}
