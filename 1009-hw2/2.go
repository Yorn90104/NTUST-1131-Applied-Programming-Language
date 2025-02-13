package main

import "fmt"

func main() {
	var (
		x      = 10
		p *int = new(int)
	)
	*p = 13
	x, *p = *p, x
	fmt.Println(x, *p)
}
