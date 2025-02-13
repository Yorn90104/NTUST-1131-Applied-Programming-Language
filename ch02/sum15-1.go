package main

import "fmt"

func IsValid(n [9]int) bool {
  var count [10]int
  for i := 0; i < len(n); i++ {
    count[n[i]] += 1
  }
  for i := 1; i <= 9; i++ {
    if count[i] != 1 {
      return false
    }
  }
  return true
}

func Solved(n [9]int) bool {
  // use every number
  if !IsValid(n) { return false }
  // row
  if n[0] + n[1] + n[2] != 15 { return false }
  if n[3] + n[4] + n[5] != 15 { return false }
  if n[6] + n[7] + n[8] != 15 { return false }
  // column
  if n[0] + n[3] + n[6] != 15 { return false }
  if n[1] + n[4] + n[7] != 15 { return false }
  if n[2] + n[5] + n[8] != 15 { return false }
  // diagnoal
  if n[0] + n[4] + n[8] != 15 { return false }
  if n[2] + n[4] + n[6] != 15 { return false }
  return true
}

func main() {
  // n[0] n[1] n[2]
  // n[3] n[4] n[5]
  // n[6] n[7] n[8]
  var n [9]int
  for i0 := 1; i0 <= 9; i0++ {
  for i1 := 1; i1 <= 9; i1++ {
  for i2 := 1; i2 <= 9; i2++ {
  for i3 := 1; i3 <= 9; i3++ {
  for i4 := 1; i4 <= 9; i4++ {
  for i5 := 1; i5 <= 9; i5++ {
  for i6 := 1; i6 <= 9; i6++ {
  for i7 := 1; i7 <= 9; i7++ {
  for i8 := 1; i8 <= 9; i8++ {
    n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8] = i0, i1, i2, i3, i4, i5, i6, i7, i8
    if Solved(n) {
      fmt.Println("Solution:")
      fmt.Println(n[0], n[1], n[2])
      fmt.Println(n[3], n[4], n[5])
      fmt.Println(n[6], n[7], n[8])
      return
    }
  }
  }
  }
  }
  }
  }
  }
  }
  }
}
