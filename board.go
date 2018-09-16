package N_Queens

import "fmt"

const N int = 4

type Point struct {
	x int
	y int
}

func CreateBoard() [][] bool {
	fmt.Println("Initializing Chess Board..")
	var board = make([][]bool, N)       // initialize a slice of dy slices
	for i:=0;i<N;i++ {
		board[i] = make([]bool, N)  // initialize a slice of dx unit8 in each of dy slices
	}

	return board
}

func PlaceQueen() {

}
