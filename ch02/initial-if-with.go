package main

import "fmt"

func main() {
  x := 3
  y := 7
  z := -2
  if p := x*y + y*z + x*z; (p * x) + ((p + z) * y) < 0 {
    fmt.Println("Error")
  } else {
    fmt.Println("OK")
  }
}
