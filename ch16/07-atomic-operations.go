package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

func next(wg *sync.WaitGroup, v *int64) {
  defer wg.Done()
  atomic.AddInt64(v, 1)
}

func main() {
  wg := &sync.WaitGroup{}
  var a int64 = 0
  wg.Add(3)
  go next(wg, &a)
  go next(wg, &a)
  go next(wg, &a)
  wg.Wait()
  fmt.Println(a)
}
