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
  fmt.Println("Is 1731 a prime number:", IsPrime(1731))
  fmt.Println("Is 3329 a prime number:", IsPrime(3329))
  fmt.Println("Is 7681 a prime number:", IsPrime(7681))
}
