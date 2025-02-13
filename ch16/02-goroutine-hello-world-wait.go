package main

import (
  "fmt"
  "sync"
)

func hello(wg *sync.WaitGroup) {
  fmt.Println("Hello World")
  wg.Done()
}

func main() {
  wg := &sync.WaitGroup{}
  wg.Add(1)
  fmt.Println("Start")
  go hello(wg)
  wg.Wait()
  fmt.Println("End")
}
