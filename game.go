package main

import(
	"time"
)

type GameState struct {
	objsys ObjectSystem
} 
var game GameState

func (s *GameState) init() {
	s.objsys.init()
	now := time.Now()
	hour := now.Hour()
	bg := &Background{file: "./res/l_night.png"}

	if hour <= 8 && hour >= 6{
		bg = &Background{file: "./res/l_sun_rise_yellow.png"}
	}else if hour <= 18 && hour >= 9{
		bg = &Background{file: "./res/l_day.png"}
	}else if hour <= 21 && hour >= 19{
		bg = &Background{file: "./res/l_sun_set.png"}
	}
	s.objsys.addObject(bg, "bg")
}

func (s *GameState) loop() {
	
}

func (s GameState) ObjectSystem() *ObjectSystem { return &s.objsys }