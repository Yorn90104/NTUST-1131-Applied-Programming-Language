package main

import "fmt"

func main() {
	count := 10
	{
		{
			count = 11
		}
		{
			count += 2
			count := 12
			count++
			fmt.Println(count)
		}
	}
	fmt.Println(count)
}
