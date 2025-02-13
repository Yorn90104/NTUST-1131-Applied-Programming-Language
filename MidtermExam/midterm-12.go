package main

// Score: 5
// Note:
// - incorrect shape of pyramid (padding of spaces is missing)
/*
(6 points)
Complete the pyramid function that prints a pyramid of height h using an
alternation of `*` and `+`. Each level in the pyramid must be annotated
with its height at the end of the level. Below are some example outputs.

Example output:
```
Height: 3
  *   0
 *+*  1
*+*+* 2
```

Example output:
```
Height: 6
     *       0
    *+*      1
   *+*+*     2
  *+*+*+*    3
 *+*+*+*+*   4
*+*+*+*+*+*  5
```

You are only allowed to:
- remove unused imported packages, and
- add code between `// -- <code begin> --` and `// -- <code end> --`.
*/

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func pyramid(h int) {
	fmt.Println("Height:", h)
	// -- <code begin> --
	for i := 0; i < h; i++ {
		fmt.Printf(strings.Repeat(" ", h-i-1))
		for j := 1; j <= (i*2)+1; j++ {
			if j%2 == 1 {
				fmt.Printf("*")
			} else {
				fmt.Printf("+")
			}
		}
		fmt.Printf(strings.Repeat(" ", h-i))
		fmt.Println(i)
	}
	// -- <code end> --
}

func main() {
	rand.Seed(time.Now().UnixNano())
	h := rand.Intn(10) + 1
	pyramid(h)
}
