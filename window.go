package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
)

var win graphics.Struct_SS_sfRenderWindow

func window_init() {
	vm := window.NewSfVideoMode()
	vm.SetWidth(800)
	vm.SetHeight(600)

	/* Create the main window */
	cs := window.NewSfContextSettings()
	win = graphics.SfRenderWindow_create(vm, "SFML window", uint(window.SfResize|window.SfClose), cs)
}