package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type MenuState struct {
	objsys ObjectSystem
} 
var menu MenuState


func (s *MenuState) init() {
	s.objsys.init()
	button := &Button{ 
		x: 200, y: 300, width: 200, height: 100, text: "test", 
		bgColor: sdl.Color{0, 255, 0, 255},
		fgColor: sdl.Color{255, 255, 255, 255}}
	button2 := &Button{ 
		x: 0, y: 0, width: 300, height: 100, text: "lol", 
		bgColor: sdl.Color{0, 0, 255, 255},
		fgColor: sdl.Color{255, 255, 255, 255}}
	text := &Text{
		x: 300, y: 200, text: "LINF", font: font, size: 100, fgColor: sdl.Color{255, 255, 255, 255}}
	s.objsys.addObject(text, "text")
	s.objsys.addObject(button, "button")
	s.objsys.addObject(button2, "button2")
	fmt.Println(button)
}

func (s *MenuState) loop() {
	
}

func (s MenuState) ObjectSystem() *ObjectSystem { return &s.objsys }
