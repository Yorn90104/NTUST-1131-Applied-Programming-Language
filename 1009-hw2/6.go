package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println("is even")
		} else {
			fmt.Println("is odd")
		}
	}
}
