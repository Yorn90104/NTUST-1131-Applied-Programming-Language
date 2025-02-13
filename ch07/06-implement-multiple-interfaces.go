package main

import "fmt"

type Speaker interface {
  Speak() string
}

type cat struct {
  name string
  age  int
}

func (c cat) String() string {
  return fmt.Sprintf("%v (%v years old)",
                     c.name,
                     c.age)
}

func (c cat) Speak() string {
  return "Meow Meow~~~"
}

func main() {
  c := cat{name: "Ore", age: 9}
  fmt.Println(c)
  fmt.Println(c.Speak())
}
