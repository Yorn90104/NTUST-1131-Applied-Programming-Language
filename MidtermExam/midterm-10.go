package main

// Score: 3
/*
(8 points)
Complete the order3 function by adding code between `// -- <code begin> --`
and `// -- <code end> --` such that the function sorts the values pointed
to by the input pointers x, y, and z in ascending order. When the function
terminates, x points to the smallest value while z points to the largest
value.

Invoke the order3 function between each pair of `// -- <invoke order3 start> --`
and `// -- <invoke order3 end> --` to sort the values of x, y, and z.

The expected output is:
```
1 3 5
2 6 8
```
*/

import "fmt"

func order3(x *int, y *int, z *int) {
	// -- <code begin> -
	if *x > *y {
		*x, *y = *y, *x
	}
	if *y > *z {
		*y, *z = *z, *y
	}
	if *x > *y {
		*x, *y = *y, *x
	}
	// -- <code end> --
}

func main() {
	x := 5
	y := 3
	z := 1
	// -- <invoke order3 start> --
	order3(&x, &y, &z)
	// -- <invoke order3 end> --
	fmt.Println(x, y, z)
	x = 8
	y = 2
	z = 6
	// -- <invoke order3 start> --
	order3(&x, &y, &z)
	// -- <invoke order3 end> --
	fmt.Println(x, y, z)
}
