package internal

import (
	"image/color"

	"github.com/erdincmutlu/board"
)

// Coordinate is used to determine piece position
type Coordinate struct {
	X int
	Y int
}

// Shape holds all different rotated views of the piece
type Shape [][]Coordinate

// Piece is to be used on the board
type Piece struct {
	Shape        Shape
	CurrentCoord Coordinate // Top left corner of the item
	Color        color.RGBA
}

type gameBoard struct {
	b *board.Board
	p *Piece
}
