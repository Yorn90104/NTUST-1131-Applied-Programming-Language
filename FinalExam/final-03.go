// Score: 10
/*
 * (10 分)
 * 解釋函式 F 的回傳值在做什麼。
 *
 * 在下面註釋內寫下你的答案（中英文皆可）：
 * -- <答案開始> --
 *  函式 F 類似一個累加器
 *  f:= F() 之後，初始化累計值 n := 0
 *  f 就變成在函式 F 中回傳的那個函式 func(m int) int
 *  每當 f(m) 時 n 就會變成 n + m  
 *  然後回傳新的 n 值
 * -- <答案結束> --
 */

package main

import "fmt"

func F() func(int) int {
  n := 0
  return func(m int) int {
    n += m
    return n
  }
}

func main() {
  f := F()
  fmt.Println(f(3))
  fmt.Println(f(5))
  fmt.Println(f(10))
}
