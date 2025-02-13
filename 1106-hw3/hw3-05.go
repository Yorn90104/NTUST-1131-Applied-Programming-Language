package main

/* (10 points)
 * Complete the function rol such that it rotates a slice
 * to the left by i items. Assume ns is a slice of integers.
 * The rotation of ns to the left by 0 item is itself. The
 * rotation of ns to the left by len(ns) items is again itself.
 * The rotation of ns to the left by i (0 < i < len(ns)) items
 * is [ns[i] ns[i+1] ... ns[len(ns)-1] ns[0] ns[1] ... ns[i-1]].
 * It is possible that the number of rotated items is larger
 * than len(ns). However, it is an error if the numer of rotated
 * items is negative. In this case, print an error message and
 * return the input slice.
 *
 * The expected output of this program is:
 *   ns                     : [1 2 3 4 5 6 7 8 9]
 *   rotate left by 3 items : [4 5 6 7 8 9 1 2 3]
 *   rotate left by 6 items : [7 8 9 1 2 3 4 5 6]
 *   rotate left by 14 items: [6 7 8 9 1 2 3 4 5]
 */

import "fmt"

func rol(ns []int, i int) []int {
	var ms []int
	// -- <code begin> --
	i %= len(ns)
	ms = append(ns[i:], ns[:i]...)
	// -- <code end> --
	return ms
}

func main() {
	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("ns                     :", ns)
	fmt.Println("rotate left by 3 items :", rol(ns, 3))
	fmt.Println("rotate left by 6 items :", rol(ns, 6))
	fmt.Println("rotate left by 14 items:", rol(ns, 14))
}
