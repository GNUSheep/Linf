package main

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Linf"
var winWidth, winHeight int32 = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	defer window.Destroy()

	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.Clear()
	defer renderer.Destroy()

	button := Button{ 
		x: 0, y: 0, width: 100, height: 100, text: "test", 
		bgColor: sdl.Color{0, 255, 0, 255},
		fgColor: sdl.Color{255, 255, 255, 255}}
	button.draw(renderer)

	gfx.CharacterColor(renderer, winWidth-16, 16, 'X', sdl.Color{255, 0, 0, 255})
	gfx.StringColor(renderer, 16, 16, "GFX Demo", sdl.Color{0, 255, 0, 255})

	renderer.Present()
	sdl.Delay(3000)

}
