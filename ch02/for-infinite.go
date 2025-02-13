package main

import "fmt"

func IsPrime(x int) bool {
  for i := 2; i < x; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  n := 10753
  for {
    n++
    if IsPrime(n) {
      break
    }
  }
  fmt.Println("The next greater prime number of 10753:", n)
}
