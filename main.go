package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
	gameloop "renfoc.us/game_loop/lib"
)

var direction = "";

type Data struct {}

func (data Data) Initialize() error {
	direction = fmt.Sprint("starting...");

	return nil;
}

func (data Data) Render() error {
	fmt.Println(direction);

	return nil;
}

func (data Data) Calculate(key keyboard.Key) error {
	switch key {
	case keyboard.KeyArrowDown:
		direction = "down"
		break;
	case keyboard.KeyArrowUp:
		direction = "up"
		break;
	case keyboard.KeyArrowLeft:
		direction = "left"
		break;
	case keyboard.KeyArrowRight:
		direction = "right"
		break;
	case keyboard.KeySpace:
		direction = "space"
		break;
	}

	return nil;
}

func main() {
	data := Data{}
	gameloop.Start(data)
}
