package internal

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
	Color              TileColor
}

// BoardPiece represent a cell in the board grid, with color
type BoardPiece struct {
	Occupied bool
	Color    TileColor
}

// Piece is a piece (to be used in draw)
type Piece struct {
	Coord Coordinate
	Color TileColor
}

// TileColor represent the color of the tile
type TileColor int

const (
	TileColorSkyBlue  TileColor = 0
	TileColorDarkBlue TileColor = 1
	TileColorOrange   TileColor = 2
	TileColorYellow   TileColor = 3
	TileColorGreen    TileColor = 4
	TileColorPurple   TileColor = 5
	TileColorRed      TileColor = 6
	TileColorSentinel TileColor = 7
)

// TileColorString is the string representation of tile colors
var TileColorString = map[TileColor]string{
	TileColorSkyBlue:  "skyblue",
	TileColorDarkBlue: "darkblue",
	TileColorOrange:   "orange",
	TileColorYellow:   "yellow",
	TileColorGreen:    "green",
	TileColorPurple:   "purple",
	TileColorRed:      "red",
}

func (c TileColor) String() string {
	return TileColorString[c]
}
