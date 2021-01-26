package main

import (
	// "github.com/veandco/go-sdl2/sdl"
)

type Element interface {
	render() bool
	handleInput() bool
}

type ObjSystem struct {
	elements []Element
}

func (o ObjSystem) addElement(e Element, name string) {
	o.elements = append(o.elements, e)
	
}
