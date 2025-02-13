package main

import (
  "fmt"
)

type Point struct {
  X, Y float64
}

func (v Point) Scale1(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v *Point) Scale2(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Point{3, 4}
  v.Scale1(10)
  fmt.Println(v.X, v.Y)
  v.Scale2(20)
  fmt.Println(v.X, v.Y)
}
