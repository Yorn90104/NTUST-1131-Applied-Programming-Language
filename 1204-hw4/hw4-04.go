/*
 * (10 points)
 * Define a variadic function named sum which accepts an arbitrary number
 * of input parameters and returns the summation of all its input parameters.
 * Assume all the input parameters are of type int. Test your implementation
 * in the main function.
 */

// -- <code begin> --
package main

import "fmt"

func sum(ns ...int) int {
	acc := 0
	for _, n := range ns {
		acc += n
	}
	return acc
}

func main() {
	fmt.Println(sum(1, 3, 5, 7, 9))
	fmt.Println(sum(2, 4, 6, 8, 10))
}

// -- <code end> --
