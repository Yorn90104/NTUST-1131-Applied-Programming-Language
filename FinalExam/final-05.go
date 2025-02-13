// Score: 40
/*
 * (50 分)
 * 實作一個排序 map ，以下是必要事項：
 * - 定義一個介面（interface）並命名為 ISortedMap ，該介面包含以下方法：
 *   + Exists：測試一個鍵（Key）是否在此排序 map 中。
 *   + Add: 將一組鍵（Key）值（Value）組合加入此排序 map 中。
 *   + Remove: 將一個鍵從此排序 map 中移除。
 *   + Keys: 將所有此排序 map 中的鍵以排序好的 slice 形式回傳。
 *   + Find: 找出一個鍵在此排序 map 中對應到的值。
 * - 定義一個結構（struct）並命名為 SortedMap ，讓 SortedMap 實作
 *   ISortedMap 。
 * - 定義一個將 ISortedMap 用作輸入參數的任意函式。
 * - 在主函式內撰寫程式碼測試你的實作。上一點的函式與所有的介面方法都要
 *   被測試到。
 *
 * 提示：
 * - comparable 型態的值可以用 == 和 != 進行比較。
 * - cmp.Ordered 型態的值可以用 < 、 <= 、 >= 和 > 進行比較。
 *
 * 評分標準：
 * - 介面（interface）定義：5 分
 * - 結構（struct）定義：5 分
 * - 方法實作之語法：5 分
 * - 方法實作之內容：
 *   + Exists：5 分
 *   + Add：5 分
 *   + Remove：5 分
 *   + Keys：5 分
 *   + Find：5 分
 * - 以 ISortedMap 作為輸入參數的函式：5 分
 * - 測試程式碼：5 分
 * - 未使用適當的型態參數（type parameters）：-2 分
 */


// -- <程式碼開始> --
package main

import (
	"fmt"
)


type ISortedMap interface{
	Exists(int) bool
	Add() 
	Remove()
	Keys()	[]int
	Find()	int
}

type SortedMap struct{
	Map map[int]int
}

func (sm SortedMap) Exists(key int)bool{
	for k := range sm.Map{
		if k == key{
			return true
		}
	}
	return false
}

// a pointer receiver is required
// the function signature is different from the signature of Add in ISortedMap
func (sm SortedMap) Add(key int, value int){
	if sm.Exists(key){
		fmt.Println("鍵已存在：",key)
	}else{
		sm.Map[key] = value
	}
}

// a pointer receiver is required
func (sm SortedMap) Remove(key int){
	if sm.Exists(key){
		delete(sm.Map, key)
	}else{
		fmt.Println("鍵不存在：",key)
	}
}

func (sm SortedMap) Keys()[]int{
	var res []int
	for k:= range sm.Map{
		res = append(res, k)
	}
        // the sorting is incomplete
	if len(res) != 1{
		for i:= 0; i <len(res) - 1; i++{
			if res[i] > res[i+1]{
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
	}
	return res
}

func (sm SortedMap) Find(key int)int{
	if sm.Exists(key){
		return sm.Map[key]
	}else{
		fmt.Println("鍵不存在：",key)
		return 0
	}
}

func main(){
	SM := SortedMap{Map:map[int]int{}}

	SM.Add(8, 7)
	SM.Add(6, 9)
	SM.Add(20, 50)
	SM.Add(200, 5)
	SM.Add(87, 69)
	SM.Add(8, 3)
	
	fmt.Println(SM.Exists(6))
	fmt.Println(SM.Exists(5))

	fmt.Println(SM.Map)
	SM.Remove(6)
	SM.Remove(9)
	fmt.Println(SM.Map)

	fmt.Println(SM.Keys())

	fmt.Println(SM.Find(8))
	fmt.Println(SM.Find(10))
}

// -- <程式碼結束> --
