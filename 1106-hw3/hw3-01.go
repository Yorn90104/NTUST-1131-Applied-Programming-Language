package main

/*
 * (10 points)
 * Complete the function appendUnique such that it appends
 * a string n to a slice of strings ns if n is not found
 * in ns.
 *
 * The expected output of this program is:
 *   [Kevin Bob Alice David]
 */

import "fmt"

func appendUnique(ns []string, n string) []string {
	// -- <code begin> --
	for _, i := range ns {
		if i == n {
			return ns
		}
	}
	ns = append(ns, n)
	// -- <code end> --
	return ns
}

func main() {
	ns := []string{"Kevin", "Bob"}
	ns = appendUnique(ns, "Alice")
	ns = appendUnique(ns, "Bob")
	ns = appendUnique(ns, "David")
	fmt.Println(ns)
}
