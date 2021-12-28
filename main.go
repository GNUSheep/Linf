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
var font_score *ttf.Font
var window *sdl.Window
var renderer *sdl.Renderer
var event sdl.Event
var statesys StateSystem
var deltaTime uint32

const fontname_score = "res/flappy-bird-font.ttf"
const fontname = "res/font.ttf"

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()
	defer sdl.Quit()

	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	defer window.Destroy()
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer renderer.Destroy()
	renderer.Present()
	var err error
	font, err = ttf.OpenFont(fontname, 200);
	if err != nil {
		fmt.Println("Font not found!!!!")
	}
	defer font.Close()
	font_score, err = ttf.OpenFont(fontname_score, 200);
	if err != nil {
		fmt.Println("Font not found!!!!")
	}
	defer font.Close()

	statesys.init(&menu)
	statesys.loop()
}
