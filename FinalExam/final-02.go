// Score: 10
/*
 * (10 分)
 * 實作一個函式，檢查兩個 map 是否相等，並在主函式內測試該函式實作。
 * 給定 m1 與 m2 這兩個 map ，如果以下兩個條件都滿足，則我們說 m1
 * 與 m2 相等：
 * - m1 與 m2 有相同長度。
 * - 對於所有 m1 裡面的鍵 k1 ， k1 也會出現在 m2 ，而且 m1[k1] == m2[k1] 。
 * 可以假設 map 的型態為 map[int]int 。但嚴禁使用 maps.Equal 。
 */


// -- <程式碼開始> --
package main

import "fmt"


func map_eql(m1 map[int]int, m2  map[int]int)bool{
	if len(m1) == len(m2){
		for k1 := range m1{
			v2, err:= m2[k1]
			if !err{
				return false
			}
			if v2 != m1[k1]{
				return false
			}
		}
		return true
	}
	return false
}

func main(){
	M1 := map[int]int{
		0:7,
		9:3,
		2:4,
	}
	M2 := map[int]int{
		0:7,
		2:4,
	}
	M3 := map[int]int{
		0:7,
		9:3,
		2:4,
	}

	fmt.Println(map_eql(M1, M2)) //false
	fmt.Println(map_eql(M1, M3)) //true
}
// -- <程式碼結束> --
