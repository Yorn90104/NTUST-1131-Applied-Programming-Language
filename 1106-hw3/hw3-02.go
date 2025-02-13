package main

/* (10 points)
 * Complete the function reversed by adding code between
 * <code begin> and <code end> such that:
 * - ms is a reverse of ns, and
 * - ns remains unchanged.
 *
 * The expected output is:
 *   ns          : [1 2 3 4 5]
 *   reversed(ns): [5 4 3 2 1]
 */

import "fmt"

func reversed(ns []int) []int {
	var ms []int
	// -- <code begin> --
	for i := len(ns) - 1; i >= 0; i-- {
		ms = append(ms, ns[i])
	}
	// -- <code end> --
	return ms
}

func main() {
	ns := []int{1, 2, 3, 4, 5}
	fmt.Println("ns          :", ns)
	fmt.Println("reversed(ns):", reversed(ns))
}
