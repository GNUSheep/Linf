package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
	"os"
)

type MenuState struct {
	objsys ObjectSystem
} 
var menu MenuState

func (s *MenuState) init() {
	s.objsys.init()
	bg := &Background{file: "./res/l_night.png", layer: 0}
	s.objsys.addObject(bg, "bg")
	for i := int32(0); i < 3; i++ {
		
		button := &Button{ 
			x: winWidth/2, y: winHeight/2+((65+10)*i)-50, width: 220, height: 65, text: " ", 
			layer: 1,
			bgColor: sdl.Color{0, 0, 255, 255},
			fgColor: sdl.Color{255, 255, 255, 255}}
		button.onClick = func() { }
		s.objsys.addObject(button, "button" + strconv.Itoa(int(i)))
	}
	s.objsys.elements["button0"].(*Button).text = "   Start   "
	s.objsys.elements["button1"].(*Button).text = "Options"
	s.objsys.elements["button2"].(*Button).text = "   Exit   "
	s.objsys.elements["button0"].(*Button).onClick = func() { statesys.setState(&game) }
	s.objsys.elements["button2"].(*Button).onClick = func() { os.Exit(0) }
	text := &Text{
		x: winWidth/2, y: winHeight/5, layer: 0, text: "LINF", font: font, size: 0.8, fgColor: sdl.Color{255, 255, 255, 255}}
	s.objsys.addObject(text, "text")

}

func (s *MenuState) loop() {

}

func (s MenuState) ObjectSystem() *ObjectSystem { return &s.objsys }
