package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/gfx"
)

const border = 5

type Button struct {
	x, y int32
	width, height int32
	text string
	bgColor sdl.Color
	fgColor sdl.Color
}

func (b Button) draw(renderer *sdl.Renderer) {
	gfx.BoxColor(renderer, b.x, b.y, b.x+b.width, b.y+b.height, b.fgColor)
	gfx.BoxColor(renderer, b.x+border, b.y+border,
		b.x+b.width-border, b.y+b.height-border, b.bgColor)
	gfx.StringColor(renderer, b.x+b.width/2, b.y+b.height/2, "GFX Demo", sdl.Color{255, 0, 0, 255})
}
