package main

import "fmt"

type WithStringName interface {
  GetName() string
}

type WithIntName interface {
  GetName() int
}

type S struct {
}

type IS S

func (s S) GetName() string {
  return "Anya"
}

func (s IS) GetName() int {
  return 7
}

func main() {
  s := S{}
  fmt.Println(s.GetName())
  fmt.Println(IS(s).GetName())
}
