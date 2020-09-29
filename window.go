package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"fmt"
)

var win *sdl.Window

func window_init() {
	var winWidth, winHeight int32 = 800, 600
	/* Create the main window */
	var err error
	win, err = sdl.CreateWindow("Linf", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(0)
	}
	//defer window.Destroy()
}