package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type OptionsState struct {
	objsys ObjectSystem
}
var options OptionsState

func (s *OptionsState) init() {
	s.objsys.init()
	options_text := &Text{
		x: winWidth/2, y: winHeight/5, layer: 0, text: "Options", font: font, size: 0.8, fgColor: sdl.Color{255, 255, 255, 255}}
	s.objsys.addObject(options_text, "options_text")

	// Controls
	options_controls := &Text{
		x: winWidth/2, y: winHeight/2 - 100, layer: 0, text: "Controls", font: font, size: 0.4, fgColor: sdl.Color{255, 255, 255, 255}}
	s.objsys.addObject(options_controls, "options_controls")
	for i := int32(0); i < 2; i++ {
		button := &Button{
			x: (winWidth/2+200) - 400 * i, y: winHeight/2 - 50, width: 220, height: 65, text: " ",
			layer: 1}
		button.onClick = func() { }
		s.objsys.addObject(button, "options_button_controls" + strconv.Itoa(int(i)))
	}
	s.objsys.elements["options_button_controls0"].(*Button).text = "Keyboard"
	s.objsys.elements["options_button_controls1"].(*Button).text = "Mouse"
	s.objsys.elements["options_button_controls0"].(*Button).onClick = func() { control_way = "keyboard" }
	s.objsys.elements["options_button_controls1"].(*Button).onClick = func() { control_way = "mouse" }

	// Difficulty
	options_level := &Text{
		x: winWidth/2, y: winHeight/2 + 75, layer: 0, text: "Difficulty", font: font, size: 0.4, fgColor: sdl.Color{255, 255, 255, 255}}
	s.objsys.addObject(options_level, "options_level")
	for i := int32(0); i < 2; i++ {
		button := &Button{
			x: (winWidth/2+200) - 400 * i, y: winHeight/2 + 150, width: 220, height: 65, text: " ",
			layer: 1}
		button.onClick = func() { }
		s.objsys.addObject(button, "options_button_level" + strconv.Itoa(int(i)))
	}
	s.objsys.elements["options_button_level0"].(*Button).text = "Normal"
	s.objsys.elements["options_button_level1"].(*Button).text = "Easy"
	s.objsys.elements["options_button_level0"].(*Button).onClick = func() { gapSize = 200 }
	s.objsys.elements["options_button_level1"].(*Button).onClick = func() { gapSize = 250 }

	// Exit
	button := &Button{
		x: 75, y: 50, width: 100, height: 65, text: " ", layer: 1}
	button.onClick = func() { }
	s.objsys.addObject(button, "options_button_exit")
	s.objsys.elements["options_button_exit"].(*Button).text = "Exit"
	s.objsys.elements["options_button_exit"].(*Button).onClick = func() { statesys.setState(&menu) }
}

func (s *OptionsState) loop() {

}

func (s OptionsState) ObjectSystem() *ObjectSystem { return &s.objsys }
