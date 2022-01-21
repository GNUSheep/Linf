package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

var Transition = false
var offset int32 = 0

type State interface {
	init()
	loop()
	ObjectSystem() *ObjectSystem
}

type StateSystem struct {
	current State
	old    State
	changed bool
	active  bool
}

func (ss *StateSystem) init(s State) {
	Transition = false
	ss.current = s
	ss.active = true
}

func (ss *StateSystem) setState(s State) {
	ss.old = ss.current
	ss.current = s
	Transition = true
	ss.changed = true
}

func (ss *StateSystem) loop() {
	for ss.active == true {
		ss.changed = false
		ss.current.init()
		if Transition {
			for i, _ := range ss.current.ObjectSystem().elements {
				if i != "bg" {
					*ss.current.ObjectSystem().elements[i].X() -= 2*winWidth
				}
			}
		}
		renderer.Present()
		if ss.changed == true {
			continue
		}
		for ss.changed == false {
			lastTime := sdl.GetTicks()
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			drawBg(renderer)
			if Transition {
				ss.old.ObjectSystem().update()
				ss.trans()
			} else {
				ss.current.loop()
			}
			ss.current.ObjectSystem().update()
			renderer.Present()
			sdl.Delay(1000/40)
			deltaTime = sdl.GetTicks() - lastTime
		}
	}
}

func (ss *StateSystem) trans() bool {
	deltaMove := int32(float32(deltaTime)/25 + 30)
	if offset < 2*winWidth {
		offset += deltaMove
	} else {
		Transition = false
		offset = 0
		return true
	}
	for i, _ := range ss.current.ObjectSystem().elements {
		if i != "bg" {
			*ss.current.ObjectSystem().elements[i].X() += deltaMove
		}
	}
	for i, _ := range ss.old.ObjectSystem().elements {
		if i != "bg" {
			*ss.old.ObjectSystem().elements[i].X() += deltaMove
		}
	}
	return false
}

func drawBg(r *sdl.Renderer) {
	now := time.Now()
	hour := now.Hour()
	bg := &Background{file: "./res/l_night.png", layer: 0}
	if hour <= 8 && hour >= 6 {
		bg = &Background{file: "./res/l_sun_rise_yellow.png", layer: 0}
	} else if hour <= 18 && hour >= 9 {
		bg = &Background{file: "./res/l_day.png", layer: 0}
	} else if hour <= 21 && hour >= 19 {
		bg = &Background{file: "./res/l_sun_set.png", layer: 0}
	}
	bg.draw(r)
}
