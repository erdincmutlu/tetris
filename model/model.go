package model

import (
	"fmt"

	"github.com/erdincmutlu/board"
)

const (
	tetrisWidth  = 10
	tetrisHeight = 20
)

type coordinate struct {
	x int
	y int
}

type shape [][]coordinate

var shapeI shape = shape{{{0, 1}, {1, 1}, {2, 1}, {3, 1}}}
var shapeReverseL shape = shape{{{0, 0}, {0, 1}, {1, 1}, {2, 1}}}
var shapeL shape = shape{{{0, 1}, {1, 1}, {2, 1}, {2, 0}}}
var shapeSq shape = shape{{{1, 0}, {1, 1}, {2, 0}, {2, 1}}}
var shapeS shape = shape{{{0, 1}, {1, 1}, {1, 0}, {2, 0}}}
var shapeT shape = shape{{{0, 1}, {1, 1}, {1, 0}, {2, 1}}}
var shapeZ shape = shape{{{0, 0}, {1, 0}, {1, 1}, {2, 1}}}

var shapesAll = []shape{shapeI, shapeReverseL, shapeL, shapeSq, shapeS, shapeT, shapeZ}

// Init will initialize the model
func Init() error {
	fmt.Println("Model init")

	_, err := board.NewBoard(tetrisWidth, tetrisHeight)
	if err != nil {
		return err
	}
	return nil
}
