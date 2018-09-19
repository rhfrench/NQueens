package main

import (
	"fmt"
	"time"
)

var startTime time.Time
var numOfScenarios int = 0
var numOfSeconds float64 = 0
var numOfSolutions int64 = 0

func main() {
	var board[][]bool = createBoard()
  	var i int = 0

	for i = 0; i <= n; i++ {
		startTime := time.Now()
		board[0][1] = true
		duration := time.Since(startTime)
		numOfSeconds = duration.Seconds()

		fmt.Println("Solutions:  ", numOfSolutions, "-- exec time:", numOfSeconds)
	}
}
