package model

import (
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

type coordinate struct {
	x int
	y int
}

// Piece is to be used on the board
type Piece struct {
	shape    shape
	location coordinate // Starting location
	color    color.RGBA
}

type shape [][]coordinate

// Tetris board 0,0 is bottom left corner
var initialLocation coordinate = coordinate{19, 3}

var shapeI shape = shape{{{0, 1}, {1, 1}, {2, 1}, {3, 1}}}
var shapeReverseL shape = shape{{{0, 0}, {0, 1}, {1, 1}, {2, 1}}}
var shapeL shape = shape{{{0, 1}, {1, 1}, {2, 1}, {2, 0}}}
var shapeSq shape = shape{{{1, 0}, {1, 1}, {2, 0}, {2, 1}}}
var shapeS shape = shape{{{0, 1}, {1, 1}, {1, 0}, {2, 0}}}
var shapeT shape = shape{{{0, 1}, {1, 1}, {1, 0}, {2, 1}}}
var shapeZ shape = shape{{{0, 0}, {1, 0}, {1, 1}, {2, 1}}}

var allShapes = []shape{shapeI, shapeReverseL, shapeL, shapeSq, shapeS, shapeT, shapeZ}
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
func NewPiece() *Piece {
	r := rand.Intn(len(allShapes))
	return &Piece{
		shape:    allShapes[r],
		location: initialLocation,
		color:    allColors[r],
	}
}

// Print will print the details of the piece
func (p *Piece) Print() {
	fmt.Printf("Piece: %+v\n", p)
}
