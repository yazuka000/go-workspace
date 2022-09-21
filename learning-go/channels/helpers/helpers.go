package helpers

import (
	"math/rand"
	"time"
)

func RandNum(n int) int {
	// rand.Seedは、乱数の元となる情報もランダム化して再定義できる
	// rand.Seedなしの場合、
	// rand.Intnは乱数生成のプログラムなのに、一番最初に出力した値しか出力しなくなる
	rand.Seed(time.Now().UnixNano())

	// valueにランダムな値を格納
	value := rand.Intn(n)
	return value
}
