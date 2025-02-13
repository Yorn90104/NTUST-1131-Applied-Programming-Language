/*
 * (10 points)
 * Define a function named countdown which accepts an input parameter
 * n of type int and returns a function of type `func() int`.
 * The returned function is served as a counter which counts from n - 1
 * to 0. Every time the returned function is called, the counter is
 * decremented by 1 (until 0) and then returned. Test your implementation
 * in the main function.
 */

// -- <code begin> --
package main

import "fmt"

func countdown(n int) func() int {
	return func() int {
		n -= 1
		return n
	}
}

func main() {
	c := countdown(10)
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

// -- <code end> --
