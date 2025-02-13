package main

/* (10 points)
 * The function insertSort has two input arguments (ns and i) and
 * one return value (ms):
 * - ns: a sorted int slice (in ascending order)
 * - i: an integer
 * - ms: also a sorted int slice (in ascending order)
 * Complete the function insertSort by adding code between
 * <code begin> and <code end> such that:
 * - ns remains unchanged
 * - ms contains exactly all the elements in ns plus i
 * - ms is sorted (in ascending order)
 *
 * The expected output is:
 *   ns               : [2 4 6 8 10]
 *   i                : 7
 *   insertSort(ns, i): [2 4 6 7 8 10]
 */

import "fmt"

func insertSort(ns []int, i int) []int {
	var ms []int
	// -- <code begin> --
	for j := 0; j < len(ns); j++ {
		ms = append(ms, ns[j])

		if j != 0 && i > ns[j] && i < ns[j+1] {
			ms = append(ms, i)
		}

	}
	// -- <code end> --
	return ms
}

func main() {
	ns := []int{2, 4, 6, 8, 10}
	i := 7
	fmt.Println("ns               :", ns)
	fmt.Println("i                :", i)
	fmt.Println("insertSort(ns, i):", insertSort(ns, i))
}
