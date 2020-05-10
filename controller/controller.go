package controller

import (
	"erdinc/tetris/model"
)

// Start will be starting point of controller
func Start() {
	piece := model.NewPiece()
	piece.Print()
}
