package main

import(
	"log"
	"Linf/mix"
)

type Sound struct{
	name *mix.Chunk
	path string
}

func (s *Sound) init(){
	s.name, _ = mix.LoadWAV(s.path)
}

func (s *Sound) play(){
	if _, err := s.name.Play(-1, 0); err != nil {
		log.Println(err)
	} 
}
