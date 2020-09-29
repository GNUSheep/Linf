package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

func menu() {
	menu_loop()
}

func menu_loop() {
	for gs == "menu" {
		menu_events()
		menu_render()
	}
}

func menu_events() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			gs = "exit"
			case *sdl.MouseMotionEvent:
			fmt.Printf("[%dms]MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
		}
	}
}
func menu_render() {
	renderer.Clear()
	renderer.Present()
	sdl.Delay(tick)
}
