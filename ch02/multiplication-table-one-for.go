package main

import "fmt"

func main() {
  for i, j := 1, 1; i <= 9; i, j = i + j / 9, j % 9 + 1 {
    fmt.Printf("%dx%d=%2d  ", j, i, i * j)
    if j == 9 {
      fmt.Println()
    }
  }
}
