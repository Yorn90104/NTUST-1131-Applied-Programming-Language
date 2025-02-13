package main

import "fmt"

type Speaker interface{
  Speak() string
}

type cat struct{
  Greeting func() string
}

func (c cat) Speak() string {
  return "Meow Meow"
}

func main() {
  c1 := cat{}
  c1.Greeting = func() string {
    return "Meow, Meow!!! mmmeeeeeeeeooooowww"
  }
  // `c1.Speak = func() string { ... }` is not allowed
  fmt.Println(c1.Speak())
  fmt.Println(c1.Greeting())

  c2 := cat{}
  c2.Greeting = func() string {
    return "mmmeeeOOWWWWW!!!"
  }
  fmt.Println(c2.Speak())
  fmt.Println(c2.Greeting())
}
