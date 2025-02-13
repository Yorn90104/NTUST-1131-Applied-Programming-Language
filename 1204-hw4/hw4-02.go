/*
 * (10 points)
 * Define a non-tail recursive function named fact which returns
 * the factorial n! of the input parameter n. Print the factorial
 * of 0, 1, 2, ..., 9 in the main function.
 */

// -- <code begin --
package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, fact(i))
	}
}

// -- <code end --
