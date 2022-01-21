package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const g = 1
const maxYSpeed = -17
const maxRot = 70

var counter = 0
var gapSize = int32(200)
var control_way = "keyboard"

type GameState struct {
	objsys ObjectSystem
}

var jumpSound = Sound{
	path: "res/jump.wav",
}

var deathSound = Sound{
	path: "res/death.wav",
}

var pointSound = Sound{
	path: "res/point.wav",
}

var game GameState

func (s GameState) ObjectSystem() *ObjectSystem { return &s.objsys }

func (s *GameState) init() {
	rand.Seed(time.Now().UnixNano())
	s.objsys.init()
	jumpSound.init()
	deathSound.init()
	pointSound.init()
	s.objsys.addObject(&Player{x: 250, y: 50, width: 50, height: 50,
		layer: 0, textureFile: "./res/bird/bird3.png"}, "player")
	var score Score
	score.init(winWidth / 2, winHeight / 5, 1)
	s.objsys.addObject(&score, "score")
	var pipes Pipes
	pipes.init(gapSize, 0, "./res/pipe.png")
	s.objsys.addObject(&pipes, "pipes")
	counter = 0
}

func (s *GameState) loop() {
	s.objsys.elements["player"].(*Player).update()
	s.objsys.elements["score"].(*Score).update()
	s.objsys.elements["pipes"].(*Pipes).update()
	checkCollisions(s.objsys.elements["pipes"].(*Pipes),
		s.objsys.elements["player"].(*Player))
}

type Player struct {
	x, y           int32
	width, height  int32
	layer          int
	textureFile    string
	accelX, accelY float32
	rotation       float64
}

func (p *Player) X() *int32   { return &p.x }
func (p *Player) Y() *int32   { return &p.y }
func (p *Player) Layer() *int { return &p.layer }

func (p *Player) update() {
	p.accelX = 0
	if p.accelY < maxYSpeed {
		p.accelY = maxYSpeed
	}
	if p.accelY > 0 {
		p.textureFile = "./res/bird/bird3.png"
	} else {
		p.textureFile = "./res/bird/bird1.png"
	}
	p.accelY += g * float32(deltaTime) / 25
	p.x += int32(p.accelX)
	p.y += int32(p.accelY)
	if maxRot > float64(p.accelY*3) {
		p.rotation = float64(p.accelY * 3)
	}
	if p.y > winHeight-p.height/2 || 0 > p.y-p.height/2 {
		statesys.setState(&over)
	}
}

func (p *Player) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture

	texfile, _ := img.Load(p.textureFile)
	texture, _ = renderer.CreateTextureFromSurface(texfile)
	defer texfile.Free()
	defer texture.Destroy()
	renderer.CopyEx(texture, nil, &sdl.Rect{p.x - (p.width / 2), p.y - (p.height / 2),
		p.width, p.height}, p.rotation, &sdl.Point{p.width / 2, p.height / 2}, 0)
}

func (p *Player) handleInput(e sdl.Event) {
	switch t := e.(type) {
	case *sdl.QuitEvent:
		os.Exit(0)
	case *sdl.KeyboardEvent:
		if t.Keysym.Sym == sdl.GetKeyFromName("Space") && t.State == 1 && control_way == "keyboard" {
			jumpSound.play()
			p.accelY += -29
		}
		if t.Keysym.Sym == sdl.K_ESCAPE {
			statesys.setState(&over)
		}
	case *sdl.MouseButtonEvent:
		if t.Button == 1 && t.State == 1 && control_way == "mouse" {
			jumpSound.play()
			p.accelY += -29
		}
	}
}

type Score struct {
	x, y  int32
	layer int
	score *Text
}

func (s *Score) X() *int32               { return &s.x }
func (s *Score) Y() *int32               { return &s.y }
func (s *Score) Layer() *int             { return &s.layer }
func (s *Score) handleInput(e sdl.Event) {}

func (s *Score) init(x int32, y int32, layer int) {
	s.x = x
	s.y = y
	s.layer = layer
	s.score = &Text{x: x, y: y, layer: layer, text: "0", font: font_score, size: 0.4, fgColor: sdl.Color{255, 255, 255, 255}}
}

var pointSoundPlayed = false

func (s *Score) update() {
	if counter >= 3 {
		s.score = &Text{
			x: s.x, y: s.y, layer: s.layer, text: strconv.Itoa(counter - 2), font: font_score, size: 0.4, fgColor: sdl.Color{255, 255, 255, 255}}
		if !pointSoundPlayed {
			pointSound.play()
			pointSoundPlayed = true
		}
	}
}

func (s *Score) draw(renderer *sdl.Renderer) {
	s.score.x = s.x
	s.score.y = s.y
	s.score.draw(renderer)
}

type Pipe struct {
	x, y        int32
	width       int32
	height      int32
	layer       int
	textureFile string
	position    string
}

func (p *Pipe) update() {
	p.x += -4
}

func (p *Pipe) X() *int32   { return &p.x }
func (p *Pipe) Y() *int32   { return &p.y }
func (p *Pipe) Layer() *int { return &p.layer }

func (p *Pipe) draw(renderer *sdl.Renderer) {
	var texture *sdl.Texture

	texfile, _ := img.Load(p.textureFile)
	texture, _ = renderer.CreateTextureFromSurface(texfile)
	defer texfile.Free()
	defer texture.Destroy()

	var width, height int32 = 100, 450
	p.width = width
	var rotation float64 = 0
	var y int32
	if p.position == "bottom" || p.position == "" {
		y = winHeight + height/2 - p.height
		rotation = 0
	} else {
		y = 0 - height/2 + p.height
		rotation = 180
	}
	renderer.CopyEx(texture, nil, &sdl.Rect{p.x - (width / 2), y - (height / 2) + p.y,
		width, height}, rotation, &sdl.Point{width / 2, height / 2}, 0)

}

func (p *Pipe) handleInput(e sdl.Event) {}

type Pipes struct {
	x, y        int32
	gapSize     int32
	layer       int
	textureFile string
	pipes       []*Pipe
	lastSpawn   time.Time
}

func (p *Pipes) update() {
	t := time.Now()
	if time.Millisecond*2500 < t.Sub(p.lastSpawn) {
		p.lastSpawn = time.Now()
		vari := int32(rand.Intn(int(winHeight/4))) - winHeight/8
		p.pipes = append(p.pipes, &Pipe{x: winWidth, y: vari, height: (winHeight / 2) - p.gapSize/2,
			layer: p.layer, position: "top", textureFile: p.textureFile})
		p.pipes = append(p.pipes, &Pipe{x: winWidth, y: vari, height: (winHeight / 2) - p.gapSize/2,
			layer: p.layer, position: "bottom", textureFile: p.textureFile})
		counter += 1
		pointSoundPlayed = false
	}
	for _, v := range p.pipes {
		v.update()
	}
	// Remove obsolete pipes
	for i, v := range p.pipes {
		if *v.X() < -v.width {
			p.pipes = append(p.pipes[:i], p.pipes[i+1:]...)
		}
	}
}

func (p *Pipes) X() *int32   { return &p.x }
func (p *Pipes) Y() *int32   { return &p.y }
func (p *Pipes) Layer() *int { return &p.layer }

func (p *Pipes) init(gapSize int32, layer int, texture string) {
	p.layer = layer
	p.gapSize = gapSize
	p.textureFile = texture
	p.lastSpawn = time.Now()
}

func (p *Pipes) draw(renderer *sdl.Renderer) {
	for _, e := range p.pipes {
		e.x += p.x
		e.draw(renderer)
		e.x -= p.x
	}
}

func (p *Pipes) handleInput(e sdl.Event) {}

func checkCollisions(pipes *Pipes, player *Player) {
	for _, v := range pipes.pipes {
		if v.x-(v.width/2) < (player.x+player.width/2) &&
			v.x+(v.width/2) > (player.x-player.width/2) {
			switch {
			case v.position == "bottom" &&
				winHeight+v.y-v.height < (player.y+player.height/2):
				statesys.setState(&over)
			case v.position == "top" &&
				v.y+v.height > (player.y-player.height/2):
				statesys.setState(&over)
			}
		}
	}
}
