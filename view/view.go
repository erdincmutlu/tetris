package view

import (
	"erdinc/tetris/internal"
	"erdinc/tetris/model"
	"image"
	"image/color"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 500
	windowHeight = 600
)

var backgroundColor color.RGBA = colornames.Darkblue

var sprite *pixel.Sprite
var win *pixelgl.Window

// Start will be starting point of view
func Start() {
	pixelgl.Run(run)
}

// For running the main window
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	var err error
	win, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tile, err := loadPicture("images/tile.png")
	if err != nil {
		panic(err)
	}

	sprite = pixel.NewSprite(tile, tile.Bounds())

	win.Clear(backgroundColor)
	border := imdraw.New(nil)
	border.Color = colornames.Red
	border.Push(pixel.V(50, windowHeight-70), pixel.V(50, windowHeight-550), pixel.V(293, windowHeight-550), pixel.V(293, windowHeight-70))
	border.Line(3)
	border.Draw(win)

	for !win.Closed() {
		drawActivePiece()
		// drawTest()
		win.Update()
	}
}

func drawPiece(coord internal.Coordinate) {
	sprite.Draw(win, pixel.IM.Moved(pixel.Vec{float64(coord.X*24 + 62), float64((19-coord.Y)*24 + 62)}))
}

func drawTest() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			drawPiece(internal.Coordinate{i, j})
		}
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func drawActivePiece() {
	activeCoords := model.GetActivePieceCoords()
	for _, coord := range activeCoords {
		drawPiece(coord)
	}
}
