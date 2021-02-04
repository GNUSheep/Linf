package main

// import (
	// "fmt"
// )

type State struct{
	init func()
	loop func()
}

type StateSystem struct {
	states []State
}
