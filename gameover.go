package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"strconv"
)

type GameOverState struct {
	objsys ObjectSystem
}
var over GameOverState

func (s *GameOverState) init() {
	s.objsys.init()
	now := time.Now()
	hour := now.Hour()
	bg := &Background{file: "./res/l_sun_set.png", layer: 0}
	if hour <= 8 && hour >= 6{
		bg = &Background{file: "./res/l_sun_rise_yellow.png", layer: 0}
	}else if hour <= 18 && hour >= 9{
		bg = &Background{file: "./res/l_day.png", layer: 0}
	}else if hour <= 21 && hour >= 19{
		bg = &Background{file: "./res/l_sun_set.png", layer: 0}
	}
	s.objsys.addObject(bg, "bg")
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
		layer: 1,
		bgColor: sdl.Color{0, 0, 255, 255},
		fgColor: sdl.Color{255, 255, 255, 255}}
	button.onClick = func() { }
	s.objsys.addObject(button, "button")
	s.objsys.elements["button"].(*Button).text = "   Play Again   "
	s.objsys.elements["button"].(*Button).onClick = func() { statesys.setState(&game) }
}

func (s *GameOverState) loop() {

}

func (s GameOverState) ObjectSystem() *ObjectSystem { return &s.objsys }
