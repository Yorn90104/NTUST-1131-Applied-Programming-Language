package main

import (
  "fmt"
  "sync"
  "time"
)

func next(wg *sync.WaitGroup, v *int, d time.Duration) {
  defer wg.Done()
  c := *v
  time.Sleep(time.Millisecond * d)
  *v = c + 1
}

func main() {
  wg := &sync.WaitGroup{}

  var d time.Duration = 0
  a := 0
  wg.Add(3)
  go next(wg, &a, d)
  go next(wg, &a, d)
  go next(wg, &a, d)
  wg.Wait()
  fmt.Printf("Sleep %d (ms): %d\n", d, a)

  d = 1
  a = 0
  wg.Add(3)
  go next(wg, &a, d)
  go next(wg, &a, d)
  go next(wg, &a, d)
  wg.Wait()
  fmt.Printf("Sleep %d (ms): %d\n", d, a)
}
