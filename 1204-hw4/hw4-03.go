/*
 * (10 points)
 * Define a tail-recursive function named fact which returns the
 * factorial n! of an input parameter n. Print the factorials of
 * 0, 1, 2, ..., and 9 in the main function.
 *
 * Hint:
 * - You may define another tail-recursive function with an extra
 *   parameter and call the function in fact.
 */

// -- <code begin --
package main

import "fmt"

func fact(n int, acc int) int {
	if n == 0 {
		return acc
	}
	return fact(n-1, n*acc)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, fact(i, 1))
	}
}

// -- <code end --
