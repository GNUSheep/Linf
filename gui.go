package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/gfx"
	sdlttf "github.com/veandco/go-sdl2/ttf"
	"fmt"
)

const border = 5
const margin = border*2

type Button struct {
	x, y int32
	width, height int32
	text string
	bgColor sdl.Color
	fgColor sdl.Color
	onClick func()
}

func (b *Button) X() *int32{ return &b.x }
func (b *Button) Y() *int32{ return &b.y }

func (b *Button) draw(renderer *sdl.Renderer) {
	gfx.BoxColor(renderer, b.x, b.y, b.x+b.width, b.y+b.height, b.fgColor)
	gfx.BoxColor(renderer, b.x+border, b.y+border,
		b.x+b.width-border, b.y+b.height-border, b.bgColor)

	var surface *sdl.Surface
	var texture *sdl.Texture

	surface, _ = font.RenderUTF8Solid(b.text, b.fgColor);
	texture, _ = renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy() 
	defer surface.Free() 

	renderer.Copy(texture, 
		&sdl.Rect{0, 0, surface.W, surface.H}, 
		&sdl.Rect{b.x+margin, b.y, b.width-(margin*2), b.height-margin})
	
}

func (b *Button) handleInput(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.MouseButtonEvent:
		fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
	}
}

type Text struct {
	x, y int32
	text string
	font *sdlttf.Font
	size int32
	fgColor sdl.Color
}

func (t *Text) X() *int32{ return &t.x }
func (t *Text) Y() *int32{ return &t.y }

func (t *Text) draw(renderer *sdl.Renderer) {
	var surface *sdl.Surface
	var texture *sdl.Texture
	surface, _ = t.font.RenderUTF8Solid(t.text, t.fgColor)
	texture, _ = renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy() 
	defer surface.Free() 
	renderer.Copy(texture, nil, &sdl.Rect{t.x, t.y, t.size, t.size})
}

func (b *Text) handleInput(sdl.Event) {}
