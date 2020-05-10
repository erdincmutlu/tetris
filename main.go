package main

import (
	"erdinc/tetris/controller"
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

	controller.Start()

	view.Start()
}
