package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	var exitcode int
	sdl.Main(func() {
		game_init()
		exitcode = game_loop()
	})
	os.Exit(exitcode)
}
