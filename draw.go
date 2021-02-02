package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	draw(*sdl.Renderer) 
	// handleInput() 
}

type DrawingSystem struct {
	elements []Drawable
}

func (o *DrawingSystem) addDrawable(e Drawable) {
	o.elements = append(o.elements, e)
}

func (o *DrawingSystem) draw(renderer *sdl.Renderer) {
	for i := 0; i < len(o.elements); i++ {
		o.elements[i].draw(renderer)
	}
}
