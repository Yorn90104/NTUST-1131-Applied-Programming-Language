package main

import "fmt"

func main() {
	for i, j := 0, 0; i <= 2 && j <= 2; i, j = (i+1)%3, j+(i+1)/3 {
		fmt.Println(i, j)
	}
}
