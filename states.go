package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type State interface {
	init()
	loop()
	ObjectSystem() *ObjectSystem
}

type StateSystem struct {
	current State
	changed bool
	active bool
}

func (ss *StateSystem) init(s State) {
	ss.current = s
	// ss.current.ObjectSystem().init()
	ss.active = true
}

func (ss *StateSystem) setState(s State) {
	ss.current = s
	// ss.current.ObjectSystem().init()
	ss.changed = true
}

func (ss *StateSystem) loop() {
	for ss.active == true {
		ss.changed = false
		ss.current.init()
		renderer.Present()
		if ss.changed == true {
			continue
		}
		for ss.changed == false {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			ss.current.loop()
			ss.current.ObjectSystem().update()
			renderer.Present()
			sdl.Delay(20)
		}
	}
}