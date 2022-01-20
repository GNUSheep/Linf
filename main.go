package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"Linf/mix"
)

var winTitle string = "Linf"
var winWidth, winHeight int32 = 800, 600
var font *ttf.Font
var btnfont *ttf.Font
var font_score *ttf.Font
var window *sdl.Window
var renderer *sdl.Renderer
var event sdl.Event
var statesys StateSystem
var deltaTime uint32

const fontname_score = "res/flappy-bird-font.ttf"
const fontname = "res/font.ttf"
const btnfontname = "res/term.ttf"

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	ttf.Init()
	defer sdl.Quit()
	mix.Init(mix.INIT_FLAC | mix.INIT_MOD | mix.INIT_MP3 | mix.INIT_OGG)
	mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE)
	defer mix.CloseAudio()
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
	btnfont, err = ttf.OpenFont(btnfontname, 12);
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
