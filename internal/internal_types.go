package internal

import (
	"image/color"

	"github.com/erdincmutlu/board"
)

type Coordinate struct {
	X int
	Y int
}

type Shape [][]Coordinate

// Piece is to be used on the board
type Piece struct {
	Shape        Shape
	CurrentCoord Coordinate
	Color        color.RGBA
}

type gameBoard struct {
	b *board.Board
	p *Piece
}
