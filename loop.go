package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
)

var eve window.SfEvent

func game_loop() {
	eve = window.NewSfEvent()
	defer window.DeleteSfEvent(eve)

	/* Start the game loop */
	for window.SfWindow_isOpen(win) > 0 {
		/* Process events */
		for window.SfWindow_pollEvent(win, eve) > 0 {
			/* Close window: exit */
			if eve.GetXtype() == window.SfEventType(window.SfEvtClosed) {
				return
			}
		}
		graphics.SfRenderWindow_clear(win, graphics.GetSfBlack())
		graphics.SfRenderWindow_display(win)
	}
}
