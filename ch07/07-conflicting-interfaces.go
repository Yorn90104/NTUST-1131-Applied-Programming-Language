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

func (s S) GetName() string {
  return "Anya"
}

func (s S) GetName() int {
  return 7
}

func main() {
  s := S{}
  fmt.Println(s.GetName())
}
