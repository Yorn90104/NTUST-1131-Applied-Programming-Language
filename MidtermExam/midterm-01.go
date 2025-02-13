package main

// Score: 0
/*
(6 points)
Modify the code such that the program outputs:
```
37
27
7
```

You are only allowed to:
- add new scopes (enclosed in `{` and `}`) and
- change `=` to `:=`.
*/

import "fmt"

// This function is used to suppress the unused variable warning.
func use(n int) {

}

func main() {
	x := 7
	y := 10
	{
		y := y + 1
		use(y)
	}
	{
		x := x + y
		use(x)

		x = x + y
		use(x)
		{
			{
				x := x + y
				use(x)
				fmt.Println(x)
			}
		}

		fmt.Println(x)
	}

	fmt.Println(x)
}
