package main

import "fmt"

type Speaker interface{
  Speak() string
}

type Cat struct{
  Speak func() string
}

func speak(s Speaker) {
  fmt.Println(s.Speak())
}


func main() {
  c := Cat{
    func () string { return "meow~~" },
  }
  speak(c)
}
