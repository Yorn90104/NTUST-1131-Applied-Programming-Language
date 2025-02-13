/*
 * (10 points)
 * Define a function `foo(n int, callback func())`. Deferred the
 * invocation of callback at the beginning of foo using a defer
 * statement. Then do anything you like in foo. Test your
 * implementation in the main function.
 */

// -- <code begin> --
package main

import "fmt"

func foo(n int, callback func()) {
	defer callback()
	fmt.Printf("老師給我 %d 分\n", n)
}

func main() {
	foo(100, func() { fmt.Println("我是 B11309044 范余振富") })
}

// -- <code end> --
