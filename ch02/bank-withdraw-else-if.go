package main

import "fmt"

func main() {
  var balance = 4958
  var withdraw = -500
  if withdraw < 0 {
    fmt.Println("Error: the withdraw amount cannot be negative")
  } else if balance >= withdraw {
    balance -= withdraw
    fmt.Println("Your new balance is", balance)
  } else {
    fmt.Println("Error: your balance is less than", withdraw)
  }
}

