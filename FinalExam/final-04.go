// Score: 9
/*
 * (10 分)
 * 建立兩個整數（Integer）通道（Channel），執行 5 個 Goroutine ，其中每一個
 * Goroutine 都會產生一個隨機整數、將該隨機整數印出、並且將該整數送至第一個
 * 通道。然後執行另一個 Goroutine 將所有從第一個通道收到的整數加總，並將最
 * 後的加總結果傳送至第二個通道。可以假設加總 Goroutine 事先知道全部會收到
 * 幾個整數。在主函式內印出從第二個通道收到的加總結果。
 */

// -- <程式碼開始> --
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
var c1 chan int
var c2 chan int

func init(){
	c1 = make(chan int)
	c2 = make(chan int)
}

func rand_num(mtx *sync.Mutex){
	mtx.Lock()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(99) +1
	fmt.Println(r)
	c1 <- r 
	mtx.Unlock()
}

func acc(){
	a := 0
	defer func() {c2 <- a}()
        // to use for range over a channel, the channel has to be closed
	for n:= range c1{
		
		a += n
	}
}

func main(){
	mtx := &sync.Mutex{}
	go rand_num(mtx)
	go rand_num(mtx)
	go rand_num(mtx)
	go rand_num(mtx)
	go rand_num(mtx)

	go acc()
	fmt.Println(<- c2)
}

// -- <程式碼結束> --
