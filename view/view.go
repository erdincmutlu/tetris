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

var backgroundColor color.RGBA = colornames.Black

var sprites [internal.TileColorSentinel]*pixel.Sprite
var border *imdraw.IMDraw
var gameOverBox *imdraw.IMDraw
var win *pixelgl.Window
var nextTxt *text.Text
var scoreTxt *text.Text
var gameOverTxt *text.Text

var score int
var gameOver bool

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

	for t := internal.TileColorSkyBlue; t < internal.TileColorSentinel; t++ {
		tile, err := loadPicture("images/tile-" + t.String() + ".png")
		if err != nil {
			panic(err)
		}
		sprites[t] = pixel.NewSprite(tile, tile.Bounds())
	}

	border = imdraw.New(nil)
	border.Color = colornames.White
	border.Push(pixel.V(50, windowHeight-70), pixel.V(50, windowHeight-550), pixel.V(293, windowHeight-550), pixel.V(293, windowHeight-70))
	border.Line(3)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	nextTxt = text.New(pixel.V(330, 500), basicAtlas)
	nextTxt.Color = colornames.White
	fmt.Fprintln(nextTxt, "Next")

	scoreTxt = text.New(pixel.V(330, 200), basicAtlas)
	scoreTxt.Color = colornames.White

	gameOverBox = imdraw.New(nil)
	gameOverBox.Color = colornames.Beige
	gameOverBox.Push(pixel.V(100, 500))
	gameOverBox.Push(pixel.V(100, 200))
	gameOverBox.Push(pixel.V(400, 200))
	gameOverBox.Push(pixel.V(400, 500))
	gameOverBox.Polygon(0)

	gameOverTxt = text.New(pixel.V(100, 400), basicAtlas)
	gameOverTxt.Color = colornames.Red
	fmt.Fprintln(gameOverTxt, "Game over\nPress enter to restart")

}

// Will initialize the window, i.e for restarting the game
func initWindow() {
	win.Clear(backgroundColor)
	border.Draw(win)
}

// Will start the handling of events loop
func startLoop() {
	last := time.Now()
	for !win.Closed() {
		dtMilliS := time.Since(last).Milliseconds()
		if gameOver {
			if win.JustPressed(pixelgl.KeyEnter) {
				gameOver = false
				score = 0
				model.NewActivePiece()
			}
		} else {
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

					completedRows := model.GetCompletedRows()
					for _, row := range completedRows {
						model.DeleteRow(row)
					}

					placed := model.NewActivePiece()
					if !(placed) {
						gameOver = true
						drawGameOver()
						model.ClearBoard()
					}
				}

				score++
				last = time.Now()
			}
		}
		win.Update()
	}
}

func drawPiece(piece internal.Piece) {
	sprites[piece.Color].Draw(win, pixel.IM.Moved(pixel.Vec{X: float64(piece.Coord.X*24 + 62), Y: float64((19-piece.Coord.Y)*24 + 62)}))
}

func drawTest() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			drawPiece(internal.Piece{Coord: internal.Coordinate{X: i, Y: j}, Color: internal.TileColor((i*10 + j) % int(internal.TileColorSentinel))})
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
	pieces := model.GetActivePieceInfo()
	for _, piece := range pieces {
		drawPiece(piece)
	}
}

// Draw the next piece to the window
func drawNextPiece() {
	pieces := model.GetNextPiece()
	for _, piece := range pieces {
		piece.Coord.X += 12
		piece.Coord.Y += 2
		drawPiece(piece)
	}
}

// Draw the board, i.e. non active pieces on the board
func drawBoard() {
	pieces := model.GetBoardPieces()
	for _, piece := range pieces {
		drawPiece(piece)
	}
}

// Draws score to the board
func drawScore() {
	scoreTxt.Clear()
	s := fmt.Sprintf("Score: %d", score)
	fmt.Fprintln(scoreTxt, s)
	scoreTxt.Draw(win, pixel.IM.Scaled(scoreTxt.Orig, 3))
}

// Draw game over message
func drawGameOver() {
	gameOverBox.Draw(win)
	gameOverTxt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(gameOverTxt.Bounds().Center())))
}
