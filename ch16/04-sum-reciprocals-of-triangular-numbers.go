package main

/*
 * Triangular numbers:
 *
 *  1   3     6       10       ...
 * -----------------------------------
 *  1   1     1       1        ...
 *     2 3   2 3     2 3
 *          4 5 6   4 5 6
 *                 7 8 9 10
 *
 * i-th (i >= 1) triangular number:
 * i * (i + 1) / 2
 *
 * Reciprocals of triangular numbers:
 * 1/1  1/3  1/6  1/10  ...
 */

import (
	"fmt"
	"sync"
	"time"
)

const (
	NUM = 10000000
)

/* i-th (i >= 1) reciprocal of triangular numbers */
func triNumRec(i int) float64 {
	return 2.0 / float64(i) / float64(i+1)
}

func sumTriNumSeq() float64 {
	sum := 0.0
	for i := 1; i <= NUM; i++ {
		sum += triNumRec(i)
	}
	return sum
}

func sumTriNumShared(tasks int) float64 {
	wg := &sync.WaitGroup{}
	wg.Add(tasks)
	sums := make([]float64, tasks)
	var sumFromTo = func(i int, j int, s int) {
		for k := i; k < j; k++ {
			sums[s] += triNumRec(k)
		}
		wg.Done()
	}
	m := NUM / tasks
	for i := range tasks {
		go sumFromTo(i*m+1, (i+1)*m, i)
	}
	wg.Wait()
	sum := 0.0
	for i := range tasks {
		sum += sums[i]
	}
	return sum
}

func sumTriNumChannel(tasks int) float64 {
	ch := make(chan float64)
	var sumFromTo = func(i int, j int) {
		sums := 0.0
		for k := i; k < j; k++ {
			sums += triNumRec(k)
		}
		ch <- sums
	}
	m := NUM / tasks
	for i := range tasks {
		go sumFromTo(i*m+1, (i+1)*m)
	}
	sums := 0.0
	for i := 0; i < tasks; i++ {
		sums += <-ch
	}
	return sums
}

func main() {
	t := time.Now()
	sum1 := sumTriNumSeq()
	d1 := time.Now().Sub(t)

	t = time.Now()
	sum2s := sumTriNumShared(2)
	d2s := time.Now().Sub(t)

	t = time.Now()
	sum4s := sumTriNumShared(4)
	d4s := time.Now().Sub(t)

	t = time.Now()
	sum2c := sumTriNumChannel(2)
	d2c := time.Now().Sub(t)

	t = time.Now()
	sum4c := sumTriNumChannel(4)
	d4c := time.Now().Sub(t)

	fmt.Println("Sequential:")
	fmt.Println("  Sum  :", sum1)
	fmt.Println("  Time :", d1)
	fmt.Println("Concurrent with shared variables (2 tasks):")
	fmt.Println("  Sum  :", sum2s)
	fmt.Println("  Time :", d2s)
	fmt.Println("Concurrent with shared variables (4 tasks):")
	fmt.Println("  Sum  :", sum4s)
	fmt.Println("  Time :", d4s)
	fmt.Println("Concurrent with channels (2 tasks):")
	fmt.Println("  Sum  :", sum2c)
	fmt.Println("  Time :", d2c)
	fmt.Println("Concurrent with channels (4 tasks):")
	fmt.Println("  Sum  :", sum4c)
	fmt.Println("  Time :", d4c)
}
