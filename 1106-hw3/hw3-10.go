/*
 * (10 points)
 * Design an exercise of Go programming. The exercise must
 * be related to what we have learnt so far. Provide your
 * suggested solution as well.

 * 製作一個"拉霸機"遊戲
 * 一局可遊玩次數：30 次
 * 通過按下 Enter 鍵進行遊戲，每按一次就會進行 1 次遊戲
 * 每次遊戲開始時，會生成三個隨機數字
 * 根據機率將隨機數歸類為不同的區間{A, B, C, D, E, F}
 *
 * 分數計算方式：
 *區間[機率]：{三個相同, 兩個相同, 不同}
 * A [36%]：{  625,   350,   150}
 * B [24%]：{ 1250,   650,   220}
 * C [17%]：{ 2100,  1080,   380}
 * D [12%]：{ 2500,  1250,   420}
 * E [ 8%]：{10000,  2500,  1250}
 * F [ 3%]：{20000, 10000,  2500}
 * 三個相同分數 = 三個相同
 * 兩個相同分數 = (兩個相同 + 不同) / 1.3
 * 三個皆不同分數 = 不同加總 / 3
 *
 * 輸出範例：
 * 機率區間： [36 60 77 89 97 100]
 * 第 1 次
 * C | B | B
 * + 790
 * 目前分數： 790
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 2 次
 * B | B | B
 * + 1250
 * 目前分數： 2040
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 3 次
 * B | A | B
 * + 610
 * 目前分數： 2650
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 4 次
 * A | B | A
 * + 430
 * 目前分數： 3080
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 5 次
 * A | E | A
 * + 1230
 * 目前分數： 4310
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 6 次
 * D | A | A
 * + 590
 * 目前分數： 4900
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 7 次
 * C | E | D
 * + 683
 * 目前分數： 5583
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 8 次
 * B | C | A
 * + 250
 * 目前分數： 5833
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 9 次
 * B | A | C
 * + 250
 * 目前分數： 6083
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 10 次
 * E | D | A
 * + 606
 * 目前分數： 6689
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 11 次
 * A | B | A
 * + 430
 * 目前分數： 7119
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 12 次
 * E | D | B
 * + 630
 * 目前分數： 7749
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 13 次
 * B | A | A
 * + 430
 * 目前分數： 8179
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 14 次
 * B | D | E
 * + 630
 * 目前分數： 8809
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 15 次
 * D | B | A
 * + 263
 * 目前分數： 9072
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 16 次
 * A | E | D
 * + 606
 * 目前分數： 9678
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 17 次
 * A | D | A
 * + 590
 * 目前分數： 10268
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 18 次
 * C | A | A
 * + 560
 * 目前分數： 10828
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 19 次
 * D | A | C
 * + 316
 * 目前分數： 11144
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 20 次
 * D | A | A
 * + 590
 * 目前分數： 11734
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 21 次
 * C | C | B
 * + 1000
 * 目前分數： 12734
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 22 次
 * B | F | D
 * + 1046
 * 目前分數： 13780
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 23 次
 * B | E | B
 * + 1460
 * 目前分數： 15240
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 24 次
 * A | C | B
 * + 250
 * 目前分數： 15490
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 25 次
 * C | A | C
 * + 940
 * 目前分數： 16430
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 26 次
 * A | C | A
 * + 560
 * 目前分數： 16990
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 27 次
 * D | A | A
 * + 590
 * 目前分數： 17580
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 28 次
 * C | C | C
 * + 2100
 * 目前分數： 19680
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 29 次
 * A | A | A
 * + 625
 * 目前分數： 20305
 * 請按ENTER

 * 機率區間： [36 60 77 89 97 100]
 * 第 30 次
 * C | A | A
 * + 560
 * 目前分數： 20865

 * 遊戲結束
 * 最終分數為：20865

 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LaBaG struct {
	//遊玩次數
	Times  int
	Played int

	//分數
	Score       int
	MarginScore int

	//內部判斷變數
	Rams   [3]int
	Prizes [3]string
	//機率與分數字典
	Rates     [6]int
	ScoreDict map[string][6]int
}

func (Game *LaBaG) result() {
	Game.Score += Game.MarginScore
	fmt.Println("第", Game.Played, "次")
	fmt.Println(Game.Prizes[0], "|", Game.Prizes[1], "|", Game.Prizes[2])
	fmt.Println("+", Game.MarginScore)
	fmt.Println("目前分數：", Game.Score)
	Game.MarginScore = 0
}

func (Game *LaBaG) GameRunning() bool {
	if Game.Played < Game.Times {
		return true
	} else {
		return false
	}
}

func (Game *LaBaG) ramdon_value() {
	//隨機數
	rand.Seed(time.Now().UnixNano())
	Game.Rams[0], Game.Rams[1], Game.Rams[2] = rand.Intn(99)+1, rand.Intn(99)+1, rand.Intn(99)+1
}

func acc_rate(Rates [6]int) []int {
	//累加機率成機率區間
	var result_rate []int
	acc := 0
	for _, Rate := range Rates {
		acc += Rate
		result_rate = append(result_rate, acc)
	}
	return result_rate
}

func (Game *LaBaG) classify_prize() {
	//分類
	rate_range := acc_rate(Game.Rates)
	fmt.Println("機率區間：", rate_range)

	for i := 0; i < 3; i++ {
		switch {
		case Game.Rams[i] <= rate_range[0]:
			Game.Prizes[i] = "A"
		case Game.Rams[i] <= rate_range[1]:
			Game.Prizes[i] = "B"
		case Game.Rams[i] <= rate_range[2]:
			Game.Prizes[i] = "C"
		case Game.Rams[i] <= rate_range[3]:
			Game.Prizes[i] = "D"
		case Game.Rams[i] <= rate_range[4]:
			Game.Prizes[i] = "E"
		case Game.Rams[i] <= rate_range[5]:
			Game.Prizes[i] = "F"
		}
	}
}

func get_score(Prize string, MarginScore int, arr [6]int) int {
	// 抓取從Array分數
	switch Prize {
	case "A":
		MarginScore += arr[0]
	case "B":
		MarginScore += arr[1]
	case "C":
		MarginScore += arr[2]
	case "D":
		MarginScore += arr[3]
	case "E":
		MarginScore += arr[4]
	case "F":
		MarginScore += arr[5]
	}
	return MarginScore
}

func (Game *LaBaG) calculate_score() {
	switch {
	case Game.Prizes[0] == Game.Prizes[1] && Game.Prizes[1] == Game.Prizes[2]: //三個相同
		Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same3"])

	case Game.Prizes[0] == Game.Prizes[1] || Game.Prizes[1] == Game.Prizes[2] || Game.Prizes[0] == Game.Prizes[2]: //兩個相同
		switch {
		case Game.Prizes[0] == Game.Prizes[1]:
			Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same2"])
			Game.MarginScore = get_score(Game.Prizes[2], Game.MarginScore, Game.ScoreDict["same1"])
		case Game.Prizes[1] == Game.Prizes[2]:
			Game.MarginScore = get_score(Game.Prizes[1], Game.MarginScore, Game.ScoreDict["same2"])
			Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same1"])
		case Game.Prizes[2] == Game.Prizes[0]:
			Game.MarginScore = get_score(Game.Prizes[2], Game.MarginScore, Game.ScoreDict["same2"])
			Game.MarginScore = get_score(Game.Prizes[1], Game.MarginScore, Game.ScoreDict["same1"])
		}
		Game.MarginScore = Game.MarginScore / 13 * 10

	case Game.Prizes[0] != Game.Prizes[1] && Game.Prizes[1] != Game.Prizes[2] && Game.Prizes[0] != Game.Prizes[2]: //三個皆不同
		Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same1"])
		Game.MarginScore = get_score(Game.Prizes[1], Game.MarginScore, Game.ScoreDict["same1"])
		Game.MarginScore = get_score(Game.Prizes[2], Game.MarginScore, Game.ScoreDict["same1"])

		Game.MarginScore /= 3
	}
}

func main() {

	Game := LaBaG{
		Times:       30,
		Played:      0,
		Score:       0,
		MarginScore: 0,
		Rams:        [3]int{0, 0, 0},
		Prizes:      [3]string{"", "", ""},
		Rates:       [6]int{36, 24, 17, 12, 8, 3},
		ScoreDict: map[string][6]int{
			"same3": {625, 1250, 2100, 2500, 10000, 20000},
			"same2": {350, 650, 1080, 1250, 5000, 10000},
			"same1": {150, 220, 380, 420, 1250, 2500},
		},
	}

	for Game.GameRunning() {
		var press string
		fmt.Println("請按ENTER")
		fmt.Scanln(&press)

		if press == "" {
			Game.ramdon_value()

			Game.classify_prize()
			Game.calculate_score()

			Game.Played++
			Game.result()
		} else {
			fmt.Println("請按ENTER")
		}

	}

	fmt.Printf("\n遊戲結束！\n最終分數為：%d", Game.Score)
}
