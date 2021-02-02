package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/gfx"
)

const border = 5
const margin = border*2

type Button struct {
	x, y int32
	width, height int32
	text string
	bgColor sdl.Color
	fgColor sdl.Color
}

func (b *Button) draw(renderer *sdl.Renderer) {
	gfx.BoxColor(renderer, b.x, b.y, b.x+b.width, b.y+b.height, b.fgColor)
	gfx.BoxColor(renderer, b.x+border, b.y+border,
		b.x+b.width-border, b.y+b.height-border, b.bgColor)

	var surface *sdl.Surface
	var texture *sdl.Texture
	surface, _ = font.RenderUTF8Solid(b.text, b.fgColor);
	texture, _ = renderer.CreateTextureFromSurface(surface)

	renderer.Copy(texture, 
		&sdl.Rect{0, 0, surface.W, surface.H}, 
		&sdl.Rect{b.x+margin, b.y, b.width-(margin*2), b.height-margin})
	
}
