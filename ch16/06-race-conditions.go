package main

import (
  "fmt"
  "sync"
)

func next(wg *sync.WaitGroup, v *int) {
  defer wg.Done()
  *v = *v + 1
}

func main() {
  wg := &sync.WaitGroup{}
  a := 0
  wg.Add(3)
  go next(wg, &a)
  go next(wg, &a)
  go next(wg, &a)
  wg.Wait()
  fmt.Println(a)
}
