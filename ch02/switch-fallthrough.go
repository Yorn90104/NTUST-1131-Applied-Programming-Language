package main

import "fmt"

func main() {
  const (
    Administrator = iota
    PowerUser
    NormalUser
    Guest
  )
  var r = NormalUser
  switch r {
    case Administrator: fallthrough
    case PowerUser: fmt.Println("Access allowed")
    case NormalUser: fallthrough
    case Guest: fallthrough
    default: fmt.Println("Access denied")
  }
}

