package main

import "runtime"

func main() {
	runtime.LockOSThread()
	window_init()
	game_loop()
}
