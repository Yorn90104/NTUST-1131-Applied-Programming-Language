package main

import (
	"fmt"
	"time"
)

var (
	day      = time.Saturday
	finished bool
	score    int = 93
)

func main() {
	fmt.Println(day, finished, score)
}
