package main

// Score: 10
/*
(10 points)
Declare variables possibly with initial values based on the following table.
Explicitly specify the variable type in a declaration if the type information
is available.

| Name | Type   | Value     |
|------|--------|-----------|
| cnt  | int    | 30        |
| pi   |        | 3.14      |
| fn   | string |           |
| chk  | *bool  | new(bool) |

After declaration, update their values according to the following table.
Here, the value of a pointer means the value pointed to by the pointer, not
the address of the pointer.

| Name | New Value  |
|------|------------|
| cnt  | cnt + 3    |
| fn   | "SayHello" |
| chk  | cnt < 100  |

After the updates, print the values of all declared variables. The expected
output is:
```
33 3.14 SayHello true
```
*/

import "fmt"

func main() {
  // -- <code begin> --
  var(
    cnt int = 30
    pi = 3.14
    fn string
    chk *bool = new(bool)
  )

  cnt += 3
  fn = "SayHello"
  *chk = cnt < 100
  fmt.Println(cnt, pi, fn, *chk)
  // -- <code end> --
}
