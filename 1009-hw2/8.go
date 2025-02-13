package main

import "fmt"

func IsPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	for n = 10753 + 2; !IsPrime(n); n += 2 {
	}
	fmt.Println("The next greater prime number of 10753 is:", n)
}
