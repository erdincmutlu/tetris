package model

import (
	"erdinc/tetris/internal"
	"fmt"
	"math/rand"
	"time"
)

const (
	tetrisWidth  = 10
	tetrisHeight = 20
)

// Tetris board 0,0 is top left corner
// 0,0 ..... 9,0
// .............
// 0,19 ...  9,19
var initialCoordinate internal.Coordinate = internal.Coordinate{X: 3, Y: 0}

// Rotation of the shapes are in defined in Super rotation system
// https://strategywiki.org/wiki/File:Tetris_rotation_super.png
var shapeI internal.Shape = internal.Shape{
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
	{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
	{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
	{{1, 0}, {1, 1}, {1, 2}, {1, 3}}}

var shapeReverseL internal.Shape = internal.Shape{
	{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
	{{1, 0}, {2, 0}, {1, 1}, {1, 2}},
	{{0, 1}, {1, 1}, {2, 1}, {2, 2}},
	{{1, 0}, {1, 1}, {1, 2}, {0, 2}}}

var shapeL internal.Shape = internal.Shape{
	{{0, 1}, {1, 1}, {2, 1}, {2, 0}},
	{{1, 0}, {1, 1}, {1, 2}, {2, 2}},
	{{0, 1}, {1, 1}, {2, 1}, {0, 2}},
	{{0, 0}, {1, 0}, {1, 1}, {1, 2}}}

var shapeSq internal.Shape = internal.Shape{
	{{1, 0}, {1, 1}, {2, 0}, {2, 1}}}

var shapeS internal.Shape = internal.Shape{
	{{0, 1}, {1, 1}, {1, 0}, {2, 0}},
	{{1, 0}, {1, 1}, {2, 1}, {2, 2}},
	{{0, 2}, {1, 2}, {1, 1}, {2, 1}},
	{{0, 0}, {0, 1}, {1, 1}, {1, 2}}}

var shapeT internal.Shape = internal.Shape{
	{{0, 1}, {1, 1}, {1, 0}, {2, 1}},
	{{1, 0}, {1, 1}, {1, 2}, {2, 1}},
	{{0, 1}, {1, 1}, {2, 1}, {1, 2}},
	{{0, 1}, {1, 0}, {1, 1}, {1, 2}}}

var shapeZ internal.Shape = internal.Shape{
	{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
	{{2, 0}, {2, 1}, {1, 1}, {1, 2}},
	{{0, 1}, {1, 1}, {1, 2}, {2, 2}},
	{{1, 0}, {1, 1}, {0, 1}, {0, 2}}}

var allShapes = []internal.Shape{shapeI, shapeReverseL, shapeL, shapeSq, shapeS, shapeT, shapeZ}
var allColors = []internal.TileColor{internal.TileColorSkyBlue, internal.TileColorDarkBlue, internal.TileColorOrange, internal.TileColorYellow, internal.TileColorGreen, internal.TileColorPurple, internal.TileColorRed}

var activePiece, nextPiece *internal.ActivePiece

var board [tetrisHeight][tetrisWidth]internal.BoardPiece

// Init will initialize the model
func Init() error {
	fmt.Println("Model init")

	rand.Seed(time.Now().Unix())
	ClearBoard()
	// RandomBoard()

	return nil
}

// ClearBoard clears the board for new game
func ClearBoard() {
	for row := 0; row < tetrisHeight; row++ {
		for col := 0; col < tetrisWidth; col++ {
			board[row][col] = internal.BoardPiece{Occupied: false}
		}
	}
}

// RandomBoard creates pseudorandomly filled board for testing
func RandomBoard() {
	for row := 10; row < 12; row++ {
		for col := 0; col < tetrisWidth-1; col++ {
			board[row][col] = internal.BoardPiece{Occupied: true}
		}
	}

	for row := 12; row < tetrisHeight; row++ {
		for col := 0; col < tetrisWidth; col++ {
			board[row][col] = internal.BoardPiece{Occupied: rand.Intn(2) != 0}
		}
	}

	board[12][9] = internal.BoardPiece{Occupied: true}
}

// NewActivePiece sets the next piece as new active piece and assign a random next piece
// return false if cannot place the active piece, i.e. game over
func NewActivePiece() bool {
	r := rand.Intn(len(allShapes))
	if nextPiece == nil {
		nextPiece = constructPiece(r)
		r = rand.Intn(len(allShapes))
	}
	activePiece = nextPiece
	nextPiece = constructPiece(r)

	if !isFit(GetActivePieceInfo()) {
		return false
	}
	return true
}

func constructPiece(r int) *internal.ActivePiece {
	return &internal.ActivePiece{
		Shape:        allShapes[r],
		CurrentCoord: initialCoordinate,
		Color:        allColors[r],
	}
}

// PrintActivePiece prints the active piece information, for debugging only
func PrintActivePiece() {
	fmt.Printf("Active Piece:%+v\n", activePiece)
}

// GetNextPiece  will return slice of pieces of the next piece
func GetNextPiece() []internal.Piece {
	return piecesFromCoords(nextPiece.Shape[0], nextPiece.Color)
}

// GetActivePieceInfo will return slice of "Piece" info for the active pieces
func GetActivePieceInfo() []internal.Piece {
	return coordsOffsetBy(activePieceCoords(), activePiece.CurrentCoord)
}

// get the current shape, i.e. slice of the coordinates, of the active piece. Shape changes as it rotates
func activePieceCoords() []internal.Piece {
	return piecesFromCoords(activePiece.Shape[activePiece.CurrentOrientation], activePiece.Color)

}

// Helper function, gets a slice of coordinates and a color, returns a slice of pieces
func piecesFromCoords(coords []internal.Coordinate, color internal.TileColor) []internal.Piece {
	var pieces []internal.Piece
	for _, coord := range coords {
		pieces = append(pieces, internal.Piece{Coord: coord, Color: color})
	}
	return pieces
}

// Returns slice of coordinates of the active piece rotated. For simulation
func getRotatedActivePieceCoords(rotateBy int) []internal.Piece {
	return coordsOffsetBy(rotatedActivePieceCoords(rotateBy), activePiece.CurrentCoord)
}

// get the rotated shape, i.e. slice of the coordinates, of the piece. Shape changes as it rotates. For simulation
func rotatedActivePieceCoords(rotateBy int) []internal.Piece {
	return piecesFromCoords(activePiece.Shape[(activePiece.CurrentOrientation+rotateBy+len(activePiece.Shape))%len(activePiece.Shape)], activePiece.Color)
}

// RotateActivePiece rotates the active piece. -1 is left, 1 is right
func RotateActivePiece(rotateBy int) {
	activePiece.CurrentOrientation = (activePiece.CurrentOrientation + rotateBy + len(activePiece.Shape)) % len(activePiece.Shape)
}

// Add offset to the given slice of pieces, to each element in the slice
func coordsOffsetBy(pieces []internal.Piece, delta internal.Coordinate) []internal.Piece {
	newCoords := make([]internal.Piece, len(pieces))
	for i := 0; i < len(pieces); i++ {
		newCoords[i] = coordOffsetBy(pieces[i], delta)
	}
	return newCoords
}

// Add offset to the given piece. Don't change the color
func coordOffsetBy(piece internal.Piece, delta internal.Coordinate) internal.Piece {
	return internal.Piece{Coord: internal.Coordinate{X: piece.Coord.X + delta.X, Y: piece.Coord.Y + delta.Y}, Color: piece.Color}
}

// CanDrop return true if the current active piece can be dropped by one
func CanDrop() bool {
	newPieces := coordsOffsetBy(GetActivePieceInfo(), internal.Coordinate{X: 0, Y: 1})
	return isFit(newPieces)
}

// Returns true if all of the pieces are in boundary and empty
func isFit(pieces []internal.Piece) bool {
	for _, piece := range pieces {
		if piece.Coord.X < 0 || piece.Coord.X >= tetrisWidth || piece.Coord.Y < 0 || piece.Coord.Y >= tetrisHeight {
			return false
		}

		if board[piece.Coord.Y][piece.Coord.X].Occupied {
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
	newCoords := coordsOffsetBy(GetActivePieceInfo(), internal.Coordinate{X: -1, Y: 0})
	return isFit(newCoords)
}

// MoveLeft will move the active piece to tne left by one
func MoveLeft() {
	activePiece.CurrentCoord.X--
}

// CanMoveRight return true if the current active piece can be moved to the right by one
func CanMoveRight() bool {
	newCoords := coordsOffsetBy(GetActivePieceInfo(), internal.Coordinate{X: 1, Y: 0})
	return isFit(newCoords)
}

// MoveRight will move the active piece to the right by one
func MoveRight() {
	activePiece.CurrentCoord.X++
}

// CanRotateLeft return true if the current active piece can be rotated left by one
func CanRotateLeft() bool {
	newCoords := getRotatedActivePieceCoords(-1)
	return isFit(newCoords)
}

// RotateLeft will rotate the active piece to tne left by one
func RotateLeft() {
	RotateActivePiece(-1)
}

// CanRotateRight return true if the current active piece can be rotated right by one
func CanRotateRight() bool {
	newCoords := getRotatedActivePieceCoords(1)
	return isFit(newCoords)
}

// RotateRight will rotate the active piece to tne right by one
func RotateRight() {
	RotateActivePiece(1)
}

// AddActivePieceToBoard adds the active piece to the board
func AddActivePieceToBoard() {
	pieces := GetActivePieceInfo()
	for _, piece := range pieces {
		board[piece.Coord.Y][piece.Coord.X].Occupied = true
		board[piece.Coord.Y][piece.Coord.X].Color = activePiece.Color
	}
}

// GetBoardPieces returns the slice of all pieces on the board
func GetBoardPieces() []internal.Piece {
	var pieces []internal.Piece
	for row := 0; row < tetrisHeight; row++ {
		for col := 0; col < tetrisWidth; col++ {
			if board[row][col].Occupied {
				pieces = append(pieces,
					internal.Piece{Coord: internal.Coordinate{X: col, Y: row}, Color: board[row][col].Color})
			}
		}
	}

	return pieces
}

// GetCompletedRows returns a slice rows ordered from highest row to lowest
func GetCompletedRows() []int {
	var completed []int

	// Only check active piece
	pieces := GetActivePieceInfo()
	for _, piece := range pieces {
		if isRowFull(piece.Coord.Y) {
			completed = addToSorted(completed, piece.Coord.Y)
		}
	}

	return completed
}

// Adds the given integer to the ordered list if item is not in already
// From lowest to highest
func addToSorted(list []int, item int) []int {
	for i := 0; i < len(list); i++ {
		if item == list[i] {
			return list
		} else if item < list[i] {
			newList := append(list[0:i], item)
			return append(newList, list[i:]...)
		}
	}

	return append(list, item)
}

// Checks if a given row is full and missing only the activeCol
func isRowFull(row int) bool {
	for col := 0; col < tetrisWidth; col++ {
		if !board[row][col].Occupied {
			return false
		}
	}

	return true
}

// DeleteRow deletes the given row from the board and drops pieces above it by one
func DeleteRow(row int) {
	if row < 0 || row >= tetrisHeight {
		return
	}

	for i := row; i > 0; i-- {
		board[i] = board[i-1]
	}
	for col := 0; col < tetrisWidth; col++ {
		board[0][col] = internal.BoardPiece{Occupied: false}
	}
}
