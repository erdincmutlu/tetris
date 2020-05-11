package model

import (
	"erdinc/tetris/internal"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"
)

const (
	tetrisWidth  = 10
	tetrisHeight = 20
)

// Tetris board 0,0 is top left corner
// 0,0 ..... 9,0
// .............
// 0,19 ...  9,19
var initialCoordinate internal.Coordinate = internal.Coordinate{3, 0}

var shapeI internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {2, 1}, {3, 1}}}
var shapeReverseL internal.Shape = internal.Shape{{{0, 0}, {0, 1}, {1, 1}, {2, 1}}}
var shapeL internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {2, 1}, {2, 0}}}
var shapeSq internal.Shape = internal.Shape{{{1, 0}, {1, 1}, {2, 0}, {2, 1}}}
var shapeS internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {1, 0}, {2, 0}}}
var shapeT internal.Shape = internal.Shape{{{0, 1}, {1, 1}, {1, 0}, {2, 1}}}
var shapeZ internal.Shape = internal.Shape{{{0, 0}, {1, 0}, {1, 1}, {2, 1}}}

var allShapes = []internal.Shape{shapeI, shapeReverseL, shapeL, shapeSq, shapeS, shapeT, shapeZ}
var allColors = []color.RGBA{colornames.Skyblue, colornames.Darkblue, colornames.Orange, colornames.Yellow, colornames.Green, colornames.Purple, colornames.Red}

var activePiece internal.Piece

var board [tetrisWidth][tetrisHeight]internal.BoardPiece

// Init will initialize the model
func Init() error {
	fmt.Println("Model init")

	rand.Seed(time.Now().Unix())
	initBoard()
	return nil
}

// Initialize the booard, to start a new game
func initBoard() {
	for i := 0; i < tetrisWidth; i++ {
		for j := 0; j < tetrisHeight; j++ {
			board[i][j] = internal.BoardPiece{Occupied: false}
		}
	}
}

// NewActivePiece sets a new random Piece as active piece
func NewActivePiece() {
	r := rand.Intn(len(allShapes))
	activePiece = internal.Piece{
		Shape:        allShapes[r],
		CurrentCoord: initialCoordinate,
		Color:        allColors[r],
	}
}

// PrintActivePiece prints the active piece information, for debugging only
func PrintActivePiece() {
	fmt.Printf("Active Piece:%+v\n", activePiece)
}

// GetActivePieceCoords will return slice of coordinates of the active piece
func GetActivePieceCoords() []internal.Coordinate {
	return coordsOffsetBy(currentActivePieceCoordinates(), activePiece.CurrentCoord)
}

// get the current shape of the piece, i.e. shape changes as it rotates
func currentActivePieceCoordinates() []internal.Coordinate {
	return activePiece.Shape[0]
}

// Add offset to the given slice of coordinates, to each element in the slice
func coordsOffsetBy(coords []internal.Coordinate, delta internal.Coordinate) []internal.Coordinate {
	newCoords := make([]internal.Coordinate, len(coords))
	for i := 0; i < len(coords); i++ {
		newCoords[i] = coordOffsetBy(coords[i], delta)
	}
	return newCoords
}

// Add offset to the given coordinate
func coordOffsetBy(coord internal.Coordinate, delta internal.Coordinate) internal.Coordinate {
	return internal.Coordinate{coord.X + delta.X, coord.Y + delta.Y}
}

// CanDrop return true if the current active piece can be dropped by one
func CanDrop() bool {
	newCoords := coordsOffsetBy(GetActivePieceCoords(), internal.Coordinate{0, 1})
	return isFit(newCoords)
}

// Returns true if all of newCoords in boundary and empty
func isFit(newCoords []internal.Coordinate) bool {
	for _, coord := range newCoords {
		if coord.X < 0 || coord.X >= tetrisWidth || coord.Y < 0 || coord.Y >= tetrisHeight {
			return false
		}

		if board[coord.X][coord.Y].Occupied {
			return false
		}
	}

	return true
}

// Drop will drop the active piece by one
func Drop() {
	activePiece.CurrentCoord.Y++
}

// CanMoveLeft return true if the current active piece can be moved to tne left by one
func CanMoveLeft() bool {
	newCoords := coordsOffsetBy(GetActivePieceCoords(), internal.Coordinate{-1, 0})
	return isFit(newCoords)
}

// MoveLeft will move the active piece to tne left by one
func MoveLeft() {
	activePiece.CurrentCoord.X--
}

// CanMoveRight return true if the current active piece can be moved to the right by one
func CanMoveRight() bool {
	newCoords := coordsOffsetBy(GetActivePieceCoords(), internal.Coordinate{1, 0})
	return isFit(newCoords)
}

// MoveRight will move the active piece to the right by one
func MoveRight() {
	activePiece.CurrentCoord.X++
}

// AddActivePieceToBoard adds the active piece to the board
func AddActivePieceToBoard() {
	coords := GetActivePieceCoords()
	for _, coord := range coords {
		board[coord.X][coord.Y].Occupied = true
		board[coord.X][coord.Y].Color = activePiece.Color
	}
}

// GetBoardPieces returns the slice of all pieces on the board
func GetBoardPieces() []internal.Coordinate {
	var pieces []internal.Coordinate
	for i := 0; i < tetrisWidth; i++ {
		for j := 0; j < tetrisHeight; j++ {
			if board[i][j].Occupied {
				pieces = append(pieces, internal.Coordinate{i, j})
			}
		}
	}

	return pieces
}
