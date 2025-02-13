package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool, 2)
	ch <- true
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
