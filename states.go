package main

import (
	"github.com/veandco/go-sdl2/sdl"
	// "fmt"
)

var Transition = false

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
			lastTime := sdl.GetTicks()
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			if Transition == true {
				// ss.old.loop()
				ss.old.ObjectSystem().update()
				ss.trans()
			} else {
				ss.current.loop()
				ss.current.ObjectSystem().update()
			}
			renderer.Present()
			sdl.Delay(1000/40)
			deltaTime = sdl.GetTicks() - lastTime
		}
	}
}

func (ss *StateSystem) trans() {
	for i, _ := range ss.old.ObjectSystem().elements {
		if i != "bg" {
			if *ss.old.ObjectSystem().elements[i].X() < 1000 && ss.old != &game {
				*ss.old.ObjectSystem().elements[i].X() += int32(float32(deltaTime)/25 + 30)
			} else {
				Transition = false
				ss.changed = true
			}
		}
	}
	// renderer.SetDrawColor(0,0,0,1)
	// rect := sdl.Rect{0,0,1000,1000}
	// renderer.DrawRect(&rect)
	// renderer.Clear()
	// renderer.Present()
	// Transition = false
	// ss.changed = true
}
