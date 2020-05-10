package view

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 500
	windowHeight = 500
)

// Start will be starting point of view
func Start() {
	pixelgl.Run(run)
}

// For running the main window
func run() {
	cfg := pixelgl.WindowConfig{
		Title: "Tetris",
		// Bounds: pixel.R(0, 0, 500, 500),
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
