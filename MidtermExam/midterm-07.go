package main

// Score: 7
// Note: the numbers start from minn, not 1
/*
(8 points)
Consider three integer variables. We define a solution as values of the
three variables that satisfy all the following conditions.
- for any integer variable n in the three variables, minn <= n <= maxn,
- the values of any two of them do not equal,
- the sum of any two of them is smaller than sum2, and
- the sum of all the three variables is greater than sum3.
Complete the function ppsols such that it prints every unique solution
exactly once in a single line (with variable values in ascending order).
Note that the order of the variable values does not matter. If `3 5 6` is
a solution, then `6 3 5` is considered as the same solution and should not
be printed (since `6 3 5` is not in ascending order).

For the test code, the expected output is:
```
-- minn = 1, maxn = 9, sum2 = 12, sum3 = 13 --
3 4 7
3 5 6
4 5 6
-- minn = 1, maxn = 9, sum2 = 10, sum3 = 11 --
3 4 5
```
*/

import "fmt"

func ppsols(minn int, maxn int, sum2 int, sum3 int) {
  // -- <code begin> --
  var n [3]int
  for n[0] = 1; minn <= n[0] && n[0] <= maxn; n[0] ++{ // the initial condition should be n[0] = minn
    for n[1] = 1; minn <= n[1] && n[1] <= maxn; n[1] ++{
      for n[2] = 1; minn <= n[2] && n[2] <= maxn; n[2] ++{
        if n[0] + n[1] + n[2] > sum3 && (n[0] + n[1] < sum2 && n[0] + n[2] < sum2 && n[1] + n[2] < sum2 ) && n[0] < n[1] && n[1] < n[2]{
          fmt.Println(n[0], n[1], n[2])
        }
      }
    }
  }
  // -- <code end> --
}

func main() {
  fmt.Println("-- minn = 1, maxn = 9, sum2 = 12, sum3 = 13 --")
  ppsols(1, 9, 12, 13)
  fmt.Println("-- minn = 1, maxn = 9, sum2 = 10, sum3 = 11 --")
  ppsols(1, 9, 10, 11)
}
