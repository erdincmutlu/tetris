package view

import (
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
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tile, err := loadPicture("images/tile.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(tile, tile.Bounds())

	border := imdraw.New(nil)

	border.Color = colornames.Red
	border.Push(pixel.V(50, windowHeight-70), pixel.V(50, windowHeight-550), pixel.V(293, windowHeight-550), pixel.V(293, windowHeight-70))
	border.Line(3)

	for !win.Closed() {
		win.Clear(backgroundColor)
		border.Draw(win)
		drawTest(sprite, win)
		win.Update()
	}
}

func drawPiece(coord []int, sprite *pixel.Sprite, win *pixelgl.Window) {
	sprite.Draw(win, pixel.IM.Moved(pixel.Vec{float64(coord[0]*24 + 62), float64(coord[1]*24 + 62)}))
}

func drawTest(sprite *pixel.Sprite, win *pixelgl.Window) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			drawPiece([]int{i, j}, sprite, win)
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
