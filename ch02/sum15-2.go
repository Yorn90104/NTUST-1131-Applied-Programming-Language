package main

import "fmt"

func IsValid(n [9]int) bool {
	var count [10]int
	for i := 0; i < len(n); i++ {
		count[n[i]] += 1
	}
	for i := 1; i <= 9; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return true
}

func Solved(n [9]int) bool {
	// use every number
	if !IsValid(n) {
		return false
	}
	// row
	if n[0]+n[1]+n[2] != 15 {
		return false
	}
	if n[3]+n[4]+n[5] != 15 {
		return false
	}
	if n[6]+n[7]+n[8] != 15 {
		return false
	}
	// column
	if n[0]+n[3]+n[6] != 15 {
		return false
	}
	if n[1]+n[4]+n[7] != 15 {
		return false
	}
	if n[2]+n[5]+n[8] != 15 {
		return false
	}
	// diagnoal
	if n[0]+n[4]+n[8] != 15 {
		return false
	}
	if n[2]+n[4]+n[6] != 15 {
		return false
	}
	return true
}

func main() {
	// n[0] n[1] n[2]
	// n[3] n[4] n[5]
	// n[6] n[7] n[8]
	var n [9]int
	var c [10]int
	for i := 0; i < len(n); i++ {
		n[i] = 1
	}
	for !Solved(n) {
		c[0] = 1
		for i := 0; i < len(n); i++ {
			c[i+1] = (n[i] + c[i]) / 10
			n[i] = (n[i] + c[i]) % 10
			if n[i] == 0 {
				n[i] = 1
			}
		}
		fmt.Println(n[0], n[1], n[2])
		fmt.Println(n[3], n[4], n[5])
		fmt.Println(n[6], n[7], n[8])
		fmt.Println("")
	}
	fmt.Println("Solution:")
	fmt.Println(n[0], n[1], n[2])
	fmt.Println(n[3], n[4], n[5])
	fmt.Println(n[6], n[7], n[8])
}
