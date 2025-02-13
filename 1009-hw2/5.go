package main

import "fmt"

func main() {

	const (
		Starfish = iota
		Titer
		Dolphin
		Cat
	)

	var (
		ani = Dolphin
	)
	switch ani {
	case 0:
		fmt.Println("Sea animal")
	case 1:
		fmt.Println("Land animal")
	case 2:
		fmt.Println("Sea animal")
	case 3:
		fmt.Println("Land animal")
	default:
		fmt.Println("Unrecognized animal")
	}
}
