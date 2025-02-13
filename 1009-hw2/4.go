package main

import "fmt"

func main() {

	var (
		name = "Kevin"
		ln   = len(name)
	)

	if ln <= 3 {
		fmt.Println("Name too short")
	} else if ln >= 10 {
		fmt.Println("Name too long")
	} else {
		fmt.Printf("The lenth or your name is:%2d.", ln)
	}
}
