package main

import (
	"fmt"
	"time"
)

func clock(ticker *time.Ticker) {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}

func main() {
	tickInterval := time.Millisecond * 500
	tickTotal := tickInterval * 3

	ticker := time.NewTicker(tickInterval)
	go clock(ticker)

	time.Sleep(tickTotal)
	ticker.Stop()

	fmt.Println("Ticker stopped")
}
