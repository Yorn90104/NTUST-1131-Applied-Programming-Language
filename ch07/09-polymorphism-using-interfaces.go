package main

import "fmt"

type Speaker interface {
  Speak() string
}

func saySomething(ss ...Speaker) {
  for _, s := range ss {
    fmt.Println(s.Speak())
  }
}

type cat struct {}
type dog struct {}
type person struct {
  name string
}

func (c cat) Speak() string {
  return "Meow Meow~~"
}

func (d dog) Speak() string {
  return "Woof Woof"
}

func (p person) Speak() string {
  return "Hi, my name is " + p.name + "."
}

func main() {
  c := cat{}
  d := dog{}
  p := person{"John"}
  saySomething(c, d, p)
}
