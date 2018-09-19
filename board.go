package main

import "fmt"

const n int = 4

type Point struct {
	x int
	y int
}

func createBoard() [][] bool {
	fmt.Println("Initializing Chess Board..")
	var board = make([][]bool, n)       // initialize a slice of dy slices
	for i:=0; i < n; i++ {
		board[i] = make([]bool, n)  // initialize a slice of dx unit8 in each of dy slices
	}

	return board
}

func PlaceQueen() {

}
