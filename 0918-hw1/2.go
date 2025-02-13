package main

import "fmt"

func main() {
	const (
		a = iota*3 + 2
		b
		c
		d
		e
	)
	fmt.Println(a, b, c, d, e)

}
