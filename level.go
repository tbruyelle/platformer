package main

import (
	"fmt"
	"github.com/tbruyelle/fsm"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/mobile/app"
)

type level struct {
	num   int
	tiles [][]rune
	objs  [][]*fsm.Object
	node  *fsm.Object
}

func loadLevel(num int) *level {
	l := &level{num: num}
	f, err := app.Open(fmt.Sprintf("%d.tiles", num))
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	l.tiles = make([][]rune, len(lines))
	for i, line := range lines {
		l.tiles[i] = make([]rune, len(line))
		for j, c := range line {
			l.tiles[i][j] = c
		}
	}
	// create the objects from the tiles
	l.node = &fsm.Object{X: 0, Y: 0, Width: float32(len(l.tiles[0])*32 + 2), Height: float32(len(l.tiles)*32 + 2)}
	l.node.Register(nil, eng)
	l.objs = make([][]*fsm.Object, len(l.tiles))
	for i, line := range l.tiles {
		l.objs[i] = make([]*fsm.Object, len(line))
		for j, tile := range line {
			o := &fsm.Object{
				X:      float32(32 * j),
				Y:      float32(32 * i),
				Width:  32,
				Height: 32,
				Sprite: tiles[tile],
				Dead:   tile == tileEmpty,
			}
			o.Register(l.node, eng)
		}
	}
	return l
}
