package view

import (
	"erdinc/tetris/internal"
	"erdinc/tetris/model"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	windowWidth  = 550
	windowHeight = 600
)

var backgroundColor color.RGBA = colornames.Darkblue

var sprite *pixel.Sprite
var win *pixelgl.Window
var nextTxt *text.Text
var scoreTxt *text.Text

var score int

// Start will be starting point of view
func Start() {
	pixelgl.Run(run)
}

// For running the main window
func run() {
	initPixel()
	initWindow()
	startLoop()
}

// Initialize pixel, sprite, etc
func initPixel() {
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

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	nextTxt = text.New(pixel.V(330, 500), basicAtlas)
	nextTxt.Color = colornames.Red
	fmt.Fprintln(nextTxt, "Next")

	scoreTxt = text.New(pixel.V(330, 200), basicAtlas)
	scoreTxt.Color = colornames.Red
}

// Will initialize the window, i.e for restarting the game
func initWindow() {
	win.Clear(backgroundColor)
	border := imdraw.New(nil)
	border.Color = colornames.Red
	border.Push(pixel.V(50, windowHeight-70), pixel.V(50, windowHeight-550), pixel.V(293, windowHeight-550), pixel.V(293, windowHeight-70))
	border.Line(3)
	border.Draw(win)
}

// Will start the handling of events loop
func startLoop() {
	last := time.Now()
	for !win.Closed() {
		dtMilliS := time.Since(last).Milliseconds()
		initWindow()
		nextTxt.Draw(win, pixel.IM.Scaled(nextTxt.Orig, 3.5))
		drawActivePiece()
		drawNextPiece()
		drawScore()
		drawBoard()

		if win.JustPressed(pixelgl.KeyLeft) {
			if model.CanMoveLeft() {
				model.MoveLeft()
			}
		} else if win.JustPressed(pixelgl.KeyRight) {
			if model.CanMoveRight() {
				model.MoveRight()
			}
		} else if win.JustPressed(pixelgl.KeyZ) {
			if model.CanRotateLeft() {
				model.RotateLeft()
			}
		} else if win.JustPressed(pixelgl.KeyX) {
			if model.CanRotateRight() {
				model.RotateRight()
			}
		}

		if dtMilliS > 1000 { // Every second, drop the active piece
			if model.CanDrop() {
				model.Drop()
			} else {
				model.AddActivePieceToBoard()
				model.NewActivePiece()
			}

			score++
			last = time.Now()
		}

		win.Update()
	}
}

func drawPiece(coord internal.Coordinate) {
	sprite.Draw(win, pixel.IM.Moved(pixel.Vec{X: float64(coord.X*24 + 62), Y: float64((19-coord.Y)*24 + 62)}))
}

func drawTest() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			drawPiece(internal.Coordinate{X: i, Y: j})
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

// Draw the active piece to the window
func drawActivePiece() {
	pieces := model.GetActivePieceCoords()
	drawCoords(pieces)
}

// Draw the next piece to the window
func drawNextPiece() {
	pieces := model.GetNextPieceCoords()
	for _, coord := range pieces {
		coord.X += 12
		coord.Y += 2
		drawPiece(coord)
	}
}

// Draw the board, i.e. non active pieces on the board
func drawBoard() {
	pieces := model.GetBoardPieces()
	drawCoords(pieces)
}

func drawCoords(pieces []internal.Coordinate) {
	for _, coord := range pieces {
		drawPiece(coord)
	}
}

func drawScore() {
	scoreTxt.Clear()
	s := fmt.Sprintf("Score: %d", score)
	fmt.Fprintln(scoreTxt, s)
	scoreTxt.Draw(win, pixel.IM.Scaled(scoreTxt.Orig, 3))
}
