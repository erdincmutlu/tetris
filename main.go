package main

import (
	"fmt"

	"github.com/erdincmutlu/board"
)

func main() {
	fmt.Println("Tetris started")
	_, err := board.NewBoard(10, 20)
	if err != nil {
		panic(err)
	}
}
