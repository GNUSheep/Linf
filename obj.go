package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Object interface {
	draw(*sdl.Renderer) 
	X() *int32
	Y() *int32
	handleInput(sdl.Event) 
}

type ObjectSystem struct {
	elements map[string]Object
}

func (o *ObjectSystem) addObject(e Object, name string) {
	o.elements[name] = e
}

func (o *ObjectSystem) init() {
	o.elements = make(map[string]Object)
}

func (o *ObjectSystem) draw(renderer *sdl.Renderer) {
	for key, _ := range o.elements {
		o.elements[key].draw(renderer)
	}
}

func (o *ObjectSystem) handleInput(e sdl.Event) {
	for key, _ := range o.elements {
		o.elements[key].handleInput(event)
	}
}
