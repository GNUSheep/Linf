package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
	"time"
)

func menu() {
	go menu_events()
	menu_loop()
}

func menu_loop() {
	for  window.SfWindow_isOpen(win) > 0 && gs == "menu" {
	}
}

func menu_events() {
	time.Sleep(2 * time.Second)
	for  window.SfWindow_isOpen(win) > 0 && gs == "menu" {
		for window.SfWindow_pollEvent(win, eve) > 0 {
			/* Close window: exit */
			if eve.GetXtype() == window.SfEventType(window.SfEvtClosed) {
				gs = "exit"
			}
			if eve.GetXtype() == window.SfEventType(window.SfEvtKeyPressed) {
				if (int(eve.GetKey().GetCode()) == window.SfKeyEscape) {
					gs = "exit"
				}
			}
		}
	}
}

func menu_render() {
		graphics.SfRenderWindow_clear(win, graphics.GetSfBlack())
		graphics.SfRenderWindow_display(win)
}
