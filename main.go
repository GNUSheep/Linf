package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var renderer *sdl.Renderer
func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, _ = sdl.CreateWindow("Linf", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	defer window.Destroy()
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.Clear()
	defer renderer.Destroy()

	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				break
			}
		}
	}
}
