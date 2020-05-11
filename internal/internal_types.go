package internal

import (
	"image/color"
)

// Coordinate is used to determine piece position
type Coordinate struct {
	X int
	Y int
}

// Shape holds all different rotated views of the piece
type Shape [][]Coordinate

// ActivePiece is to be used on the board
type ActivePiece struct {
	Shape              Shape
	CurrentCoord       Coordinate // Top left corner of the item
	CurrentOrientation int        // Orientation changes by rotating left or right
	Color              color.RGBA
}

// BoardPiece represent a cell in the board grid, with color
type BoardPiece struct {
	Occupied bool
	Color    color.RGBA
}
