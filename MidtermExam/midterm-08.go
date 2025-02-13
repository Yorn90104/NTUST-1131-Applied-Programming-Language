package main

// Score: 3
// Note:
// - not in-place reversal
// - incorrect initialization of slices
/*
(6 points)
Complete the reverse function that inplace reverses a slice of integers.
The expected output is:
```
[100]
[9 7 1 3 5]
[11 7 3 19 17 13]
```

You are only allowed to add code between `// -- <code begin> --` and
`// -- <code end> --`.

Hints:
- For a slice ns, len(ns) is the length (the number of elements) of the slice.
- For a slice ns, ns[i] denotes the i-th (starting from 0) element in the slice.
- For example, if ns = []int{5, 3, 1, 7, 9}, then
  + ns[0] is 5 and ns[1] is 3, and
  + With ns[1] = 11, the value of ns becomes []int{5, 11, 1, 7, 9}
*/

import "fmt"

func reverse(ns []int) {
	// -- <code begin> --
	l := len(ns)
	for i := 0; i*2 <= l; i++ {
		ns[i], ns[l-i-1] = ns[l-i-1], ns[i]
	}
	// -- <code end> --
}

func main() {
	ns1 := []int{100}
	ns2 := []int{5, 3, 1, 7, 9}
	ns3 := []int{13, 17, 19, 3, 7, 11}
	reverse(ns1)
	reverse(ns2)
	reverse(ns3)
	fmt.Println(ns1)
	fmt.Println(ns2)
	fmt.Println(ns3)
}
