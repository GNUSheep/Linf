package main

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
	ss.active = true
}

func (ss *StateSystem) setState(s State) {
	ss.current = s
	ss.changed = true
}

func (ss *StateSystem) loop() {
	for ss.active == true {
		ss.changed = false
		ss.current.init()
		if ss.changed == true {
			continue
		}
		for ss.changed == false {
			ss.current.loop()
			ss.current.ObjectSystem().handleInput(event)
			ss.current.ObjectSystem().draw(renderer)
		}
	}
}
