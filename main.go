package main

import (
	"fmt"

	"github.com/erdincmutlu/board"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	fmt.Println("Tetris started")
	_, err := board.NewBoard(tetrisWidth, tetrisHeight)
	if err != nil {
		panic(err)
	}

	pixelgl.Run(run)

}

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

	win.Clear(colornames.Darkblue)

	for !win.Closed() {
		win.Update()
	}
}
