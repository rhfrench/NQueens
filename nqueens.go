package main

import (
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	x int
	y int
}

var boardResults = make([][]Point, 0)

func main() {
	Solve(4)

}

func Solve(n int) {
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		CheckRecursively(start, current, n)
	}
	fmt.Print("Results:\n\n")

	var resultCount int = 0

	for _, result := range boardResults {
		resultCount++
		fmt.Println("Result # " + strconv.Itoa(resultCount))

		for _, resultQueenLocation := range result {

			fmt.Print(resultQueenLocation)
		}
		fmt.Print("\n")
	}

	fmt.Printf("Total Results: %d", len(boardResults))
}
func CheckRecursively(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			boardResults = append(boardResults, c)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					CheckRecursively(nextStart, current, n)
				}
			}
		}
	}
}
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	//fmt.Print(a, b)
	answer := a.x == b.x || a.y == b.y || math.Abs(float64(a.y-b.y)) == math.Abs(float64(a.x-b.x))
	//fmt.Print(answer)
	return answer
}
