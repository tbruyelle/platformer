package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/tbruyelle/fsm"

	"golang.org/x/mobile/app"
)

const (
	tileFloor = '1' + iota
	tilePlain
	tileWallR
	tileWallE
	tileCeiling
	tileBall
	tileMax

	tilePlayer = 'P'
)

type level struct {
	fsm.Object
	num              int
	tiles            [][]rune
	objs             [][]*fsm.Object
	playerX, playerY float32
	maxX, maxY       float32
}

func loadLevel(num int) *level {
	l := &level{
		num: num,
		Object: fsm.Object{
			X:      0,
			Y:      0,
			Width:  1,
			Height: 1,
		},
	}

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
			if c == tilePlayer {
				// Found where the player starts
				player = &fsm.Object{
					X:      screenHalfW,
					Y:      screenHalfH,
					Width:  32,
					Height: 32,
					Sprite: tiles[tileBall],
				}
				player.Register(scene, eng)
				l.playerX, l.playerY = float32(j*32), float32(i*32)
			}
		}
	}
	l.maxX = float32(len(l.tiles[0]) * 32)
	l.maxY = float32(len(l.tiles) * 32)
	l.playerMove()
	log.Println("level max", l.maxX, l.maxY)

	// create the objects from the tiles
	l.Register(scene, eng)
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
				Dead:   tile >= tileMax,
			}
			o.Register(&l.Object, eng)
		}
	}
	return l
}

func (l *level) playerMove() {
	x := l.playerX
	y := l.playerY

	switch {
	case x < screenHalfW:
		l.X = 0
	case x > l.maxX-screenHalfW:
		l.X = -l.maxX + screenW
	default:
		l.X = -x + screenHalfW
	}
	switch {
	case y < screenHalfH:
		l.Y = 0
	case y > l.maxY-screenHalfH:
		l.Y = -l.maxY + screenH
	default:
		l.Y = -y + screenHalfH
	}
}
