package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n1 := rand.Intn(999) + 1
	n2 := rand.Intn(999) + 1
	fmt.Printf("GCD of %d and %d is ", n1, n2)
	for n2 > 0 {
		n1, n2 = n2, n1%n2
	}
	fmt.Printf("%d\n", n1)
}
