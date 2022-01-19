package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"sort"
	//"fmt"
	//"time"
)

type Object interface {
	draw(*sdl.Renderer) 
	X() *int32
	Y() *int32
	Layer() *int
	handleInput(sdl.Event) 
}

type ObjectSystem struct {
	elements map[string]Object
	layers map[int][]Object
}

func (o *ObjectSystem) addObject(e Object, name string) {
	o.elements[name] = e
	o.layers[*e.Layer()] = append(o.layers[*e.Layer()], e)
}

func (o *ObjectSystem) init() {
	o.elements = make(map[string]Object)
	o.layers = make(map[int][]Object)
}

func (o *ObjectSystem) draw(renderer *sdl.Renderer) {
	keys := make([]int, 0, len(o.layers))
	for k, _ := range o.layers {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for key, _ := range keys {
		for num, _ := range o.layers[key] {
			o.layers[key][num].draw(renderer)
		}
	}
}

func (o *ObjectSystem) handleInput() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		for key, _ := range o.elements {
			o.elements[key].handleInput(event)
		}
	}
}

func (o *ObjectSystem) update() {
	o.handleInput()
	o.draw(renderer)
	if Transition == true {
		o.Show()
	}
}

func (o *ObjectSystem) Show(){
		for i, _ := range o.elements {
			if i != "bg"{
				if *o.elements[i].X() < 1000{
					*o.elements[i].X() += int32(float32(deltaTime)/25 + 30)
				}else{
					statesys.setState(&game)
				}
			}
		}
}
