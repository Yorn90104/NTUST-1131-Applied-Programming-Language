// Score: 9
/*
 * (10 分)
 * 實作一個函式，將兩個已排序（以某種遞增順序）整數（Integer）slice 合併為
 * 另一個已排序 slice （以相同的順序），並在在主函式（Main Function）中測試
 * 該函式實現。
 *
 * 範例:
 * - 合併 [1 4 8 9] 和 [3 8 10] 的結果為： [1 3 4 8 8 9 10]。
 * - 合併 [7 8 11] 和 [2 5 9 10] 的結果為： [2 5 7 8 9 10 11]。
 *
 * 提示（假設函式的兩個輸入參數為 ns 與 ms ）：
 * - 使用兩個變數 i 與 j 記住目前瀏覽 ns 與 ms 這兩個 slice 的位置。
 * - 在一個 for 迴圈內，總是比對 ns[i] 與 ms[j] 的大小，並且將兩者之中較小
 *   的數附加（append）在結果之後。
 * - 根據 ns[i] 與 ms[j] 的大小，將 i 或 j 增加 1 。
 * - 當 for 迴圈完整看完其中一個 slice 時，迴圈結束。要注意此時另一個 slice
 *   還有未被瀏覽過的項目需要被附加於結果之後。
 *
 * 注意：
 * - 嚴禁使用任何外部排序函式（例如 sort.Slice）。
 */

// -- <程式碼開始> --
package main

import "fmt"

func merge(ns []int, ms []int) []int {
	var res []int
	
	i, j := 0, 0
	for !(i == len(ns)- 1 && j == len(ms) - 1){
		if ns[i] > ms[j]{
			res = append(res, ms[j])
                        // j should always be advanced by 1 here
			if j < len(ms) - 1{
				j++
			}
			
		}else{
			res = append(res, ns[i])
                        // i should always be advanced by 1 here
			if i < len(ns) - 1{
				i++
			}
			
		}
	}
        // there may be many items left to be appended
	if ns[i] > ms[j]{
		res = append(res, ms[j])
		res = append(res, ns[i])
	}else{
		res = append(res, ns[i])
		res = append(res, ms[j])
	}
	return res
}

func main(){
	ns := []int{1, 4, 8, 9}
	ms := []int{3, 8, 10}
	fmt.Println(merge(ns, ms))
	ns = []int{7, 8, 11}
	ms = []int{2, 5, 9, 10}
	fmt.Println(merge(ns, ms))
}


// -- <程式碼結束> --
