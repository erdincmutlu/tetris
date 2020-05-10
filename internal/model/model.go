package model

import (
	"erdinc/tetris/internal"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/erdincmutlu/board"
	"golang.org/x/image/colornames"
)

const (
	tetrisWidth  = 10
	tetrisHeight = 20
)

// Tetris board 0,0 is bottom left corner
// 0,19 ...  9,19
// .............
// 0,0 ..... 9,0
var initialCoordinate internal.Coordinate = internal.Coordinate{19, 3}

var shapeI internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {2, 1}, {3, 1}}}
var shapeReverseL internal.Shape = internal.Shape{{{0, 0}, {0, 1}, {1, 1}, {2, 1}}}
var shapeL internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {2, 1}, {2, 0}}}
var shapeSq internal.Shape = internal.Shape{{{1, 0}, {1, 1}, {2, 0}, {2, 1}}}
var shapeS internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {1, 0}, {2, 0}}}
var shapeT internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {1, 0}, {2, 1}}}
var shapeZ internal.Shape = internal.Shape{{{0, 0}, {1, 0}, {1, 1}, {2, 1}}}

var allShapes = []internal.Shape{shapeI, shapeReverseL, shapeL, shapeSq, shapeS, shapeT, shapeZ}
var allColors = []color.RGBA{colornames.Skyblue, colornames.Darkblue, colornames.Orange, colornames.Yellow, colornames.Green, colornames.Purple, colornames.Red}

// Init will initialize the model
func Init() error {
	fmt.Println("Model init")

	rand.Seed(time.Now().Unix())
	_, err := board.NewBoard(tetrisWidth, tetrisHeight)
	if err != nil {
		return err
	}
	return nil
}

// NewPiece returns a new random Piece
func NewPiece() *internal.Piece {
	r := rand.Intn(len(allShapes))
	return &internal.Piece{
		Shape:        allShapes[r],
		CurrentCoord: initialCoordinate,
		Color:        allColors[r],
	}
}
