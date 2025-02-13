package main

import "fmt"

type cat struct {
  name string
  age  int
}

func (c cat) String() string {
  return fmt.Sprintf("%v (%v years old)",
                     c.name,
                     c.age)
}

func main() {
  c := cat{name: "Ore", age: 9}
  fmt.Println(c)
}
