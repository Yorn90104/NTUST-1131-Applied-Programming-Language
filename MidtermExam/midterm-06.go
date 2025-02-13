package main

// Score: 10
/*
(10 points)
Follow the instructions below to complete the function factorial such that
it prints n! (the factorial of n) if n is non-negative; otherwise it prints
the error message "The input must be non-negative".
- Add an if statement that checks if n is smaller than 0.
  + In the body of the if statement, print the error message and return.
- Declare variable res with initial value 1.
- Add a for i loop.
  + Initial statement of the for loop: declare i with initial value 2.
  + Condition of the for loop: i is smaller than or equal to n.
  + Post statement of the for loop: i is incremented by 1.
  + Body of the for loop: multiply res by i and store the multiplication back to res.
- Print the value of res.

The expected output of the test code is:
```
120
The input must be non-negative
```
*/

import "fmt"

func factorial(n int) {
  // -- <code begin> --
  if n < 0{
    fmt.Println("The input must be non-negative")
  }else {
    res := 1
    for i := 2; i <= n ; i++ {
      res *= i
    }
    fmt.Println(res)
  }
  
  // -- <code end> --
}

func main() {
  factorial(5)
  factorial(-3)
}
