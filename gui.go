package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/gfx"
	sdlttf "github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/img"
	"fmt"
)

const border = 5
const margin = border*2

type Button struct {
	x, y int32
	layer int
	width, height int32
	text string
	bgColor sdl.Color
	fgColor sdl.Color
	onClick func()
}

func (b *Button) X() *int32{ return &b.x }
func (b *Button) Y() *int32{ return &b.y }
func (b *Button) Layer() *int{ return &b.layer }

func (b *Button) draw(renderer *sdl.Renderer) {
	gfx.BoxColor(renderer, b.x-(b.width/2), b.y-(b.height/2), b.x+(b.width/2), b.y+(b.height/2), b.bgColor)
	// I dont like the borders xd
	// gfx.BoxColor(renderer, b.x+border, b.y+border,
		// b.x+b.width-border, b.y+b.height-border, b.bgColor)

	var surface *sdl.Surface
	var texture *sdl.Texture

	surface, _ = font.RenderUTF8Solid(b.text, b.fgColor);
	texture, _ = renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy() 
	defer surface.Free() 

	renderer.Copy(texture, 
		&sdl.Rect{0, 0, surface.W, surface.H}, 
		&sdl.Rect{b.x-(b.width/2)+margin, b.y-(b.height/2)+margin, b.width-(margin*2), b.height-2*margin})
	
}

func (b *Button) handleInput(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.MouseButtonEvent:
		// fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
		// t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
		if b.x-(b.width/2) < t.X && t.X < b.x+(b.width/2) &&
		   b.y-(b.height/2) < t.Y && t.Y < b.y+(b.height/2) {
			switch t.State {
			case 1:
				fmt.Println(b)
				b.bgColor.A = 150
			case 0:
				b.onClick()
				b.bgColor.A = 255
			}
		}
	}
}

type Text struct {
	x, y int32
	layer int
	text string
	font *sdlttf.Font
	size float32
	fgColor sdl.Color
}

func (t *Text) X() *int32{ return &t.x }
func (t *Text) Y() *int32{ return &t.y }
func (t *Text) Layer() *int{ return &t.layer }

func (t *Text) draw(renderer *sdl.Renderer) {
	var surface *sdl.Surface
	var texture *sdl.Texture
	surface, _ = t.font.RenderUTF8Solid(t.text, t.fgColor)
	texture, _ = renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy() 
	defer surface.Free() 
	width := int32(float32(surface.W)*t.size)
	height := int32(float32(surface.H)*t.size)
	renderer.Copy(texture, nil, &sdl.Rect{t.x-(width/2), t.y-(height/2), width, height})
}

func (b *Text) handleInput(sdl.Event) {}

type Background struct {
	x, y int32
	layer int
	file string
}

func (bg *Background) X() *int32{ return &bg.x }
func (bg *Background) Y() *int32{ return &bg.y }
func (bg *Background) Layer() *int{ return &bg.layer }

func (bg *Background) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture

	bgfile, _ := img.Load(bg.file)
	texture, _ = renderer.CreateTextureFromSurface(bgfile)
	defer bgfile.Free()
	defer texture.Destroy()

	renderer.Copy(texture, nil, &sdl.Rect{0, 0, 800, 600})
}

func (bg *Background) handleInput(sdl.Event) {}

