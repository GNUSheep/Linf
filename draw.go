package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type Drawable interface {
	draw(*sdl.Renderer) 
	X() *int32
	Y() *int32
	// handleInput() 
}

type DrawingSystem struct {
	elements map[string]Drawable
}

func (o *DrawingSystem) addDrawable(e Drawable, name string) {
	o.elements[name] = e
}

func (o *DrawingSystem) init() {
	o.elements = make(map[string]Drawable)
}

func (o *DrawingSystem) draw(renderer *sdl.Renderer) {
	// for i := 0; i < len(o.elements); i++ {
		// o.elements[i].draw(renderer)
	// }
	for key, value := range o.elements {
		fmt.Println("Key:", key, "Value:", value)
		o.elements[key].draw(renderer)
	}
}
