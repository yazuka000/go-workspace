package main

import (
	"github.com/yazuka000/myniceprogram/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandNum(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	// ゴルーチン
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
