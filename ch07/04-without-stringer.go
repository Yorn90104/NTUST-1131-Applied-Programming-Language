package main

import "fmt"

type cat struct {
  name string
  age  int
}

func main() {
  c := cat{name: "Ore", age: 9}
  fmt.Println(c)
}
