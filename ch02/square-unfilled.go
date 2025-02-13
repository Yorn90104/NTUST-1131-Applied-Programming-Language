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
  s1 := strings.Repeat("*", n)
  s2 := ""
  if n >= 2 {
    s2 = "*" + strings.Repeat(" ", n - 2) + "*"
  }
  fmt.Println("n =", n)
  for i := 0; i < n; i++ {
    if i == 0 || i == n - 1 {
      fmt.Println(s1)
    } else {
      fmt.Println(s2)
    }
  }
}
