package main

import "fmt"

func main() {
	var x, y, z int = 3, 7, 11
	x, y, y = x+y, y+z, x+z
	fmt.Println(x, y, z)
}
