package main

import (
	"fmt"
)

var ch chan int
var done chan bool

func init() {
	ch = make(chan int, 1)
	done = make(chan bool, 0)
}

func producer(n int) {
	defer close(ch)
	for i := 0; i < n; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
	}
}

func consumer() {
	defer func() { done <- true }()
	for n := range ch {
		fmt.Printf("Received: %d\n", n)
	}
}

func main() {
	go consumer()
	go producer(7)
	<-done
}
