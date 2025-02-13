package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	var Height int
	rand.Seed(time.Now().UnixNano())

	Height = rand.Intn(10) + 1
	fmt.Println("Height: ", Height)

	for i := 1; i <= Height; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
	for j := Height - 1; j > 0; j-- {
		fmt.Println(strings.Repeat("*", j))
	}
}
