package main

import (
	"fmt"
)

func fib(i int) func(yield func(int) bool) {
	curr, next := 0, 1
	return func(yield func(int) bool) {
		for _ = range 10 {
			if !yield(curr) {
				return
			}
			curr, next = next, curr+next
		}
	}
}

func main() {

	fmt.Println(fib(10)(1))

}
