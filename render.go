package main

import (
	"fmt"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
)

func render() {
	for window.SfWindow_isOpen(win) > 0 {
		switch gs {
		case "menu":
			menu_render()
		case "playing":
			playing_render()
		case "exit":
			continue
		default:
			fmt.Println("Render error")
		}
	}
}
