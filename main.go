package main

import (
	// "github.com/veandco/go-sdl2/gfx"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var winTitle string = "Linf"
var winWidth, winHeight int32 = 800, 600
var font *ttf.Font
var window *sdl.Window
var renderer *sdl.Renderer
var event sdl.Event

const fontname = "res/noto.ttf"

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()
	defer sdl.Quit()

	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	defer window.Destroy()

	var err error

	font, err = ttf.OpenFont(fontname, 200);
	if err != nil {
		fmt.Println("Font not found!1!!")
	}
	defer font.Close()

	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.Clear()
	defer renderer.Destroy()
	renderer.Present()

	var drawsys ObjectSystem
	drawsys.init()
	button := &Button{ 
		x: 200, y: 300, width: 200, height: 100, text: "test", 
		bgColor: sdl.Color{0, 255, 0, 255},
		fgColor: sdl.Color{255, 255, 255, 255}}
	button2 := &Button{ 
		x: 0, y: 0, width: 300, height: 100, text: "lol", 
		bgColor: sdl.Color{0, 255, 0, 255},
		fgColor: sdl.Color{255, 255, 255, 255}}
	text := &Text{
		x: 300, y: 200, text: "LINF", font: font, size: 100, fgColor: sdl.Color{255, 255, 255, 255}}
	drawsys.addObject(text, "text")
	drawsys.addObject(button, "button")
	drawsys.addObject(button2, "button2")
	drawsys.draw(renderer)
	renderer.Present()
	sdl.Delay(3000)
	*drawsys.elements["button"].X() = 400
	drawsys.draw(renderer)

	renderer.Present()
	sdl.Delay(30000)
}
