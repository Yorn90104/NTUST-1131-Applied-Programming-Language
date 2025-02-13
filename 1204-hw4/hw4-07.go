/*
 * (10 points)
 * Define a function named mkslice which returns a slice filled
 * with some identical item. The item and the size of the returned
 * slice are given as input parameters of mkslice. The type of the
 * returned slice depends on the type of the given item. For example,
 * mkslice("a", 4) should return the slice [a a a a] of type []string,
 * and mkslice(2, 5) should return the slice [2 2 2 2 2] of type []int.
 * Use type parameters in mkslice. Test your implementation in the
 * main function.
 */

// -- <code begin> --
package main

import "fmt"

func mkslice[T any](item T, n int) []T {
	var res []T
	for i := 0; i < n; i++ {
		res = append(res, item)
	}
	return res
}

func main() {
	fmt.Println(mkslice("a", 4))
	fmt.Println(mkslice(2, 5))
}
