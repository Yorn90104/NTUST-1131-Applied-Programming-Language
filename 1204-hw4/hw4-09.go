/*
 * (10 points)
 * Write a program that continually asks the user to input an integer
 * until the user inputs "q". Print the summation of all the integers
 * input by the user.
 *
 * Hint:
 * - Use fmt.Scan(&s) to ask for an input and store the input in
 *   a pre-declared variable s
 * - Use strconv.Atoi to convert a string to an int
 */

// -- <code begin> --
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		s  string
		ns []int
	)
	for true {
		fmt.Println("請輸入數字，不再輸入請輸入'q'：")
		fmt.Scanln(&s)
		if s != "q" {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("錯誤的輸入！", err)
				os.Exit(1)
			}
			ns = append(ns, n)
		} else {
			break
		}
	}
	res := 0
	for i := range ns {
		res += ns[i]
		if i != len(ns)-1 {
			fmt.Printf("%d +", ns[i])
		} else {
			fmt.Printf("%d", ns[i])
		}
	}
	fmt.Println("=", res)
}

// -- <code end> --
