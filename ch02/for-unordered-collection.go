package main

import "fmt"

func main() {
  items := []string{ "Charger", "Laptop", "Tablet" }
  for i, item := range items {
    fmt.Printf("The %d-th item is: %s\n", i, item)
  }
}
