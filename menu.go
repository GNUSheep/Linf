package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	"fmt"
	"os"
)

var menu_bg_image *sdl.Surface
var menu_bg_image_path string = "./res/l_day.png"
var menu_bg_texture *sdl.Texture
var src,dst sdl.Rect = sdl.Rect{0, 0, 999, 999}, sdl.Rect{0, 0, winWidth, winHeight}

func menu() {
	var err error
	menu_bg_image, err = img.Load(menu_bg_image_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(0)
	}
	menu_bg_texture, err = renderer.CreateTextureFromSurface(menu_bg_image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(0)
	}
	menu_loop()
}

func menu_loop() {
	for gs == "menu" {
		menu_events()
		menu_render()
	}
}

func menu_events() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			gs = "exit"
			case *sdl.MouseMotionEvent:
			fmt.Printf("[%dms]MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
		}
	}
}
func menu_render() {
	renderer.Clear()
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{0, 0, int32(winWidth), int32(winHeight)})
	renderer.Copy(menu_bg_texture, &src, &dst)
	renderer.Present()
	sdl.Delay(tick)
}
