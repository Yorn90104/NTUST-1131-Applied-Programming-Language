package main

// Score: 8
/*
(8 points)
Complete the until711 function such that the function visits integers
n+1, n+2, n+3, ... before reaching an integer that is both a multiple of
7 and 11. The function prints "Seven" if the visited integer is a multiple
of 7 while it prints "Eleven" if the visited integer is a multiple of 11.
The variable cnt is used to count the number of visited 7's and 11's
multiples. Update cnt correctly in your code.

For n = 37, the expected output is:
```
42 Seven
44 Eleven
49 Seven
55 Eleven
56 Seven
63 Seven
66 Eleven
70 Seven
Multiples visited: 8
```

For n = 78, the expected output is:
```
84 Seven
88 Eleven
91 Seven
98 Seven
99 Eleven
105 Seven
110 Eleven
112 Seven
119 Seven
121 Eleven
126 Seven
132 Eleven
133 Seven
140 Seven
143 Eleven
147 Seven
Multiples visited: 16
```

You are only allowed to add code between `// -- <code begin> --` and
`// -- <code end> ---`.
*/

import "fmt"

func until711(n int) {
  i := n + 1
  cnt := 0
  // -- <code begin> --
  for { 
    if i % 7 == 0 && i % 11 == 0{
      break
    }else if i % 7 == 0{
      fmt.Println(i, "Seven")
      cnt += 1
    }else if i % 11 == 0{
      fmt.Println(i, "Eleven")
      cnt += 1
    }
    i++
  }
  // -- <code end> --
  fmt.Println("Multiples visited:", cnt)
}

func main() {
  until711(37)
  fmt.Println("--")
  until711(78)
}
