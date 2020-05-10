package main

import (
	"erdinc/tetris/internal/controller"
	"erdinc/tetris/internal/model"
	"erdinc/tetris/internal/view"

	"fmt"
)

func main() {
	fmt.Println("Tetris started")

	// model.Hello()

	err := model.Init()
	if err != nil {
		panic(err)
	}

	controller.Start()
	// controller.Start()

	view.Start()
}
