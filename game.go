package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"fmt"
)

// Global sfEvent varible
var event sdl.Event
// Global game stage varible (can be set to "menu" or "playing")
var gs string
var renderer *sdl.Renderer
var tick uint32

func game_init() {
	runtime.LockOSThread()
	window_init()
	var err error
	renderer, _ = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(0)
	}

	/* Select the game stage */
	gs = "menu"
	tick = 20
}

func game_loop() int {
	/* Start the game loop */
	for true {
		/* Game stage selection loop */
		/* If a stage exits the next one is choosen */
		switch gs {
		case "menu":
			menu()
		// case "playing":
			// playing()
		case "exit":
			return 0
		default:
			return 0
		}
	}
	return 0
}



