package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/window"
	"time"
	"os"
)

// Global sfEvent varible
var eve window.SfEvent
// Global game stage varible (can be set to "menu" or "playing")
var gs string

var ticker *time.Ticker

func game_init() {
	window_init()
	/* Initialize event handler */
	eve = window.NewSfEvent()
	defer window.DeleteSfEvent(eve)

	/* Select the game stage */
	gs = "menu"
	ticker = time.NewTicker(20 * time.Millisecond)
}

func game_loop() {
	/* Start the render loop */
	go render()

	/* Start the game loop */
	for window.SfWindow_isOpen(win) > 0 {
		/* Game stage selection loop */
		/* If a stage exits the next one is choosen */
		switch gs {
		case "menu":
			menu()
		case "playing":
			playing()
		case "exit":
			os.Exit(0)
		default:
			os.Exit(0)
		}
	}
}



