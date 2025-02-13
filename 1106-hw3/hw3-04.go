package main

/* (10 points)
 * Complete the function interleave by adding code between
 * <code begin> and <code end> such that:
 * - ns and ms remain unchanged,
 * - rs is the interleaving of ns and ms, that is
 *   + if len(ns) >= len(ms), then rs is
 *     [ns[0], ms[0], ns[1], ms[1], ..., ns[len(ms)-1], ms[len(ms)-1], ns[len(ms)], ..., ns[len(ns)-1]
 *   + if len(ms) >= len(ns), then rs is
 *     [ns[0], ms[0], ns[1], ms[1], ..., ns[len(ns)-1], ms[len(ns)-1], ns[len(ns)], ..., ns[len(ms)-1]
 *
 * The expected output is:
 *   ns          : [1 3 5 7 9]
 *   ms          : [2 4 6 8 10]
 *   interleaving: [1 2 3 4 5 6 7 8 9 10]
 *   --
 *   ns          : [6 2 4]
 *   ms          : [1 3 10 5 7]
 *   interleaving: [6 1 2 3 4 10 5 7]
 */

import "fmt"

func interleave(ns []int, ms []int) []int {
	var rs []int
	// -- <code begin> --
	for i := 0; i < len(ns) || i < len(ms); i++ {
		if i < len(ns) {
			rs = append(rs, ns[i])
		}
		if i < len(ms) {
			rs = append(rs, ms[i])
		}
	}
	// -- <code end> --
	return rs
}

func main() {
	ns := []int{1, 3, 5, 7, 9}
	ms := []int{2, 4, 6, 8, 10}
	fmt.Println("ns          :", ns)
	fmt.Println("ms          :", ms)
	fmt.Println("interleaving:", interleave(ns, ms))
	fmt.Println("--")
	ns = []int{6, 2, 4}
	ms = []int{1, 3, 10, 5, 7}
	fmt.Println("ns          :", ns)
	fmt.Println("ms          :", ms)
	fmt.Println("interleaving:", interleave(ns, ms))
}
