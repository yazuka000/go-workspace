package main

import (
	"github.com/yazuka000/myniceprogram/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandNum(numPool)

	// チャネルに値を送信
	intChan <- randomNumber
}

func main() {
	// 新しいチャネルを作成
	intChan := make(chan int, 1)

	// deferで呼び出した関数は、どこに定義されても一番最後に呼び出される。つまりfinally
	// 下の例は、チャネルintChanをcloseし、値を送受信しないようにすることで、余計な誤作動を避けている
	defer close(intChan)

	// ゴルーチン
	go CalculateValue(intChan)

	// チャネルにストックされた値をnumに格納
	num := <-intChan
	log.Println(num)
}
