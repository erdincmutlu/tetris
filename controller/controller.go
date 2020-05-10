package controller

import (
	"erdinc/tetris/model"
	"fmt"
)

// Start will be starting point of controller
func Start() {
	piece := model.NewPiece()
	fmt.Printf("Piece: %+v\n", piece)
}
