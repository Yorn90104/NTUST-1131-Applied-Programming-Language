package main

// Score: 4
// Note:
// - each rune has to be converted back to string for printing
/*
(8 points)
Add code between `// -- <code begin> --` and `// -- <code end> --` such that
the function pstrsp prints the input string str with additional "_" between
every pair of consecutive characters.

The expected output of this program is:
```
P_e_t_e_r
B_ü_c_h_i
イ_ル_カ_は_と_て_も_か_わ_い_い_で_す
```
*/

import "fmt"

func pstrsp(str string) {
	// -- <code begin> --
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]))
		if i < len(runes)-1 {
			fmt.Print("_")
		}
	}
	// -- <code end> --
	fmt.Println()
}

func main() {
	pstrsp("Peter")
	pstrsp("Büchi")
	pstrsp("イルカはとてもかわいいです")
}
