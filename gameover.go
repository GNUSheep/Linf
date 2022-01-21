package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type GameOverState struct {
	objsys ObjectSystem
}
var over GameOverState

func (s *GameOverState) init() {
	s.objsys.init()
	gameover_text := &Text{
		x: winWidth/2, y: winHeight/5, layer: 0, text: "Game Over", font: font, size: 0.8, fgColor: sdl.Color{255, 255, 255, 255}}
	s.objsys.addObject(gameover_text, "gameover_text")
	if counter < 3 {
		gameover_score := &Text{
			x: winWidth/2, y: winHeight/5 + 150, layer: 0, text: strconv.Itoa(0), font: font_score, size: 0.4, fgColor: sdl.Color{255, 255, 255, 255}}
		s.objsys.addObject(gameover_score, "gameover_score")
	}else{
		gameover_score := &Text{
			x: winWidth/2, y: winHeight/5 + 150, layer: 0, text: strconv.Itoa(counter - 2), font: font_score, size: 0.4, fgColor: sdl.Color{255, 255, 255, 255}}
		s.objsys.addObject(gameover_score, "gameover_score")
	}
	button := &Button{
		x: winWidth/2, y: winHeight/5 + 250, width: 220, height: 65, text: " ",
		layer: 1}
	button.onClick = func() { }
	s.objsys.addObject(button, "button")
	s.objsys.elements["button"].(*Button).text = "   Play Again   "
	s.objsys.elements["button"].(*Button).onClick = func() { statesys.setState(&menu) }
	deathSound.play()
}

func (s *GameOverState) loop() {

}

func (s GameOverState) ObjectSystem() *ObjectSystem { return &s.objsys }
