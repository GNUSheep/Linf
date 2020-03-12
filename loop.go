package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/window"
)

// Global sfEvent varible
var eve window.SfEvent
// Global game stage varible (can be set to "menu" or "playing")
var gs string

func game_loop() {
	eve = window.NewSfEvent()
	defer window.DeleteSfEvent(eve)
	gs = "menu"

	/* Start the render loop */
	go render()
	
	/* Start the game loop */
	for window.SfWindow_isOpen(win) > 0 {
		/* Process events */
		switch gs {
		case "menu":
			menu()
		case "playing":
			playing()
		default:
			return
		}
	}
}



