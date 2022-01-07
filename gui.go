package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	sdlttf "github.com/veandco/go-sdl2/ttf"
)

const border = 5
const margin = border * 2

type Button struct {
	x, y          int32
	layer         int
	width, height int32
	text          string
	isPressed     bool
	onClick       func()
}

func (b *Button) X() *int32   { return &b.x }
func (b *Button) Y() *int32   { return &b.y }
func (b *Button) Layer() *int { return &b.layer }

func (b *Button) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture
	var bgTexture string
	var colorDelta uint8
	if b.isPressed {
		bgTexture = "./res/button_pressed.png"
		colorDelta = 100
	} else {
		bgTexture = "./res/button.png"
		colorDelta = 0
	}
	texfile, _ := img.Load(bgTexture)
	texture, _ = renderer.CreateTextureFromSurface(texfile)
	defer texfile.Free()
	defer texture.Destroy()

	renderer.Copy(texture, nil,
		&sdl.Rect{b.x - (b.width / 2), b.y - (b.height / 2), (b.width), (b.height)})

	surface, _ := btnfont.RenderUTF8Solid(b.text,
		sdl.Color{255 - colorDelta, 255 - colorDelta, 255 - colorDelta, 255})
	text_texture, _ := renderer.CreateTextureFromSurface(surface)
	defer text_texture.Destroy()
	defer surface.Free()

	prop := float64(surface.W) / float64(surface.H)
	text_height := b.height - (2 * margin) - 1
	text_width := int32(float64(text_height)*prop) - (margin * 2)
	xoff := text_width / 2

	renderer.Copy(text_texture,
		&sdl.Rect{0, 0, surface.W, surface.H},
		&sdl.Rect{b.x - xoff, b.y - (b.height / 2) + margin, text_width, text_height})
}

func (b *Button) handleInput(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.MouseButtonEvent:
		if b.x-(b.width/2) < t.X && t.X < b.x+(b.width/2) &&
			b.y-(b.height/2) < t.Y && t.Y < b.y+(b.height/2) {
			switch t.State {
			case 1:
				b.isPressed = true
			case 0:
				b.onClick()
				b.isPressed = false
			}
		}
	}
}

type Text struct {
	x, y    int32
	layer   int
	text    string
	font    *sdlttf.Font
	size    float32
	fgColor sdl.Color
}

func (t *Text) X() *int32   { return &t.x }
func (t *Text) Y() *int32   { return &t.y }
func (t *Text) Layer() *int { return &t.layer }

func (t *Text) draw(renderer *sdl.Renderer) {
	var surface *sdl.Surface
	var texture *sdl.Texture
	surface, _ = t.font.RenderUTF8Solid(t.text, t.fgColor)
	texture, _ = renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy()
	defer surface.Free()
	width := int32(float32(surface.W) * t.size)
	height := int32(float32(surface.H) * t.size)
	renderer.Copy(texture, nil, &sdl.Rect{t.x - (width / 2), t.y - (height / 2), width, height})
}

func (t *Text) handleInput(sdl.Event) {}

type Background struct {
	layer int
	file  string
}

func (bg *Background) X() *int32   { return nil }
func (bg *Background) Y() *int32   { return nil }
func (bg *Background) Layer() *int { return &bg.layer }

func (bg *Background) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture
	bgfile, _ := img.Load(bg.file)
	texture, _ = renderer.CreateTextureFromSurface(bgfile)
	defer bgfile.Free()
	defer texture.Destroy()
	renderer.Copy(texture, nil, &sdl.Rect{0, 0, 800, 600})
}

func (bg *Background) handleInput(sdl.Event) {}
