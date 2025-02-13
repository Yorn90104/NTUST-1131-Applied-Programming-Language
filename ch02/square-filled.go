package main

import (
  "fmt"
  "strings"
  "time"
  "math/rand"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  n := rand.Intn(9) + 1
  s := strings.Repeat("*", n)
  fmt.Println("n =", n)
  for i := 0; i < n; i++ {
    fmt.Println(s)
  }
}
