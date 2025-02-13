/*
 * (10 points)
 * Perform the following tasks the main function.
 * - Assign an anonymous function to a variable. The anonymous
 *   Function simply prints "Hello World".
 * - Invoke the function through the variable.
 */

// -- <code begin> --
package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Hello World")
	}
	f()
}

// -- <code end> --
