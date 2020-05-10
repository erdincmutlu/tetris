package controller

import (
	"erdinc/tetris/model"
)

// Start will be starting point of controller
func Start() {
	model.NewActivePiece()
	model.PrintActivePiece()
}
