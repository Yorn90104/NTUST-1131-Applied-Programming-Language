package main

import "fmt"

func main() {
  x := 3
  y := 7
  z := -2
  if ((x*y + y*z + x*z) * x) + ((x*y + y*z + x*z + z) * y) < 0 {
    fmt.Println("Error")
  } else {
    fmt.Println("OK")
  }
}
