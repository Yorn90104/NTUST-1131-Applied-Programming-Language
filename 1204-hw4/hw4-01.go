/*
 * (10 points)
 * A Fibonacci sequence is series of numbers F(0), F(1), F(2), ... where
 * - F(0) = 0
 * - F(1) = 1
 * - F(n) = F(n - 1) + F(n - 2), for n > 1
 * Define a recursive function named fib which returns F(i) for a given
 * input parameter i. Any loop is forbidden. Print the first ten numbers
 * in Fibonacci sequence in the main function.
 */

// -- <code begin --

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
	for i := range fib(10) {
		fmt.Printf("%d ", i)
	}
}

// -- <code end --
