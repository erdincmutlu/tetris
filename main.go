package main

import (
	"erdinc/tetris/model"
	"erdinc/tetris/view"
	"fmt"
)

func main() {
	fmt.Println("Tetris started")
	err := model.Init()
	if err != nil {
		panic(err)
	}

	view.Start()
}
