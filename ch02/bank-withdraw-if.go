package main

import "fmt"

func main() {
  var balance = 4958
  var withdraw = 5000
  if balance >= withdraw {
    balance -= withdraw
    fmt.Println("Your new balance is", balance)
  }
  if balance < withdraw {
    fmt.Println("Error: your balance is less than", withdraw)
  }
}

