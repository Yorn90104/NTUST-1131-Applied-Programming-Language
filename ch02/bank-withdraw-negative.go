package main

import "fmt"

func main() {
  var balance = 4958
  var withdraw = -500
  if balance >= withdraw {
    balance -= withdraw
    fmt.Println("Your new balance is", balance)
  } else {
    fmt.Println("Error: your balance is less than", withdraw)
  }
}

