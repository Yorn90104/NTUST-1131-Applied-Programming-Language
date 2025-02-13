package main

// Score: 2
// Note: incorrect string literals
/*
(10 points)
Declare two string variables rstr and istr between `// -- <code begin> --` and
`// -- <code end> --`.  Also initialize rstr using a raw string literal and
initialize istr using an interpreted string literal such that this program prints:
```
-- Raw --
WSL can be found at "C:\Program Files\WSL"
-- Interpreted --
Code is enclosed in backticks (`).
For example: Use `code` in a Markdown file.
```
*/

import "fmt"

func main() {
	// -- <code begin> --
	var (
		rstr string = `WSL can be found at "C:\Program Files\WSL"`
		istr string = "Code is enclosed in backticks (`).\nFor example: Use `code` in a Markdown file."
	)
	// -- <code end> --
	fmt.Println("-- Raw --")
	fmt.Println(rstr)
	fmt.Println("-- Interpreted --")
	fmt.Println(istr)
}
