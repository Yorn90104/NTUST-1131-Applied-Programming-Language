package main

import (
	"fmt"
	"sync"
)

func next(wg *sync.WaitGroup, mtx *sync.Mutex, v *int64) {
	defer wg.Done()
	mtx.Lock()
	c := *v
	*v = c + 1
	mtx.Unlock()
}

func main() {
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	var a int64 = 0
	wg.Add(3)
	go next(wg, mtx, &a)
	go next(wg, mtx, &a)
	go next(wg, mtx, &a)
	wg.Wait()
	fmt.Println(a)
}
