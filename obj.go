package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type Object interface {
	draw(*sdl.Renderer) 
	X() *int32
	Y() *int32
	handleInput() 
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
	// for i := 0; i < len(o.elements); i++ {
		// o.elements[i].draw(renderer)
	// }
	for key, value := range o.elements {
		fmt.Println("Key:", key, "Value:", value)
		o.elements[key].draw(renderer)
	}
}
