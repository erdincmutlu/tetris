package model

import (
	"fmt"

	"github.com/erdincmutlu/board"
)

const (
	tetrisWidth  = 10
	tetrisHeight = 20
)

// Init will initialize the model
func Init() error {
	fmt.Println("Model init")

	_, err := board.NewBoard(tetrisWidth, tetrisHeight)
	if err != nil {
		return err
	}
	return nil
}
