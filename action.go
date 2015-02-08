package main

import (
	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/sprite/clock"
	"log"
)

type moveTo struct {
	x, y float32
	v    *fsm.Vector
}

const (
	playerSpeed = 10
)

func (a *moveTo) Do(o *fsm.Object, t clock.Time) {
	if o.Time == 0 {
		o.Time = t
		// Fix target x,y to match player x,y
		a.x, a.y = a.x-player.Width/2, a.y-player.Height/2
		// Compute vector to direction
		a.v = fsm.NewVector(player.X, player.Y, a.x, a.y)
		// Normalize vector
		a.v.Normalize()
		// Turn target x,y to level plan
		a.x, a.y = a.x-lvl.X, a.y-lvl.Y
		return
	}
	// Check if the goal is reached
	v := fsm.NewVector(player.X-lvl.X, player.Y-lvl.Y, a.x, a.y)
	l := v.Length()
	log.Println("lvl", lvl.X, lvl.Y)
	log.Println("player", player.X, player.Y)
	log.Println("v", v.X, a.v.Y, l)
	log.Println("-------")
	if l <= playerSpeed {
		// move over
		log.Println("move over")
		o.Reset()
		o.Action = nil
		return
	}
	f := clock.EaseIn(o.Time, o.Time+10, t)
	vx, vy := a.v.X*playerSpeed*f, a.v.Y*playerSpeed*f
	scroll(vx, vy)
}

func scroll(vx, vy float32) {
	// Where to apply those values ?
	switch {
	case lvl.X-vx > 0 || player.X < screenHalfW:
		// level coordinates should never be positives
		player.X = player.X - lvl.X + vx
		lvl.X = 0
	case lvl.X-vx < -lvl.maxX:
		// should never be lower than the level maxs
		player.X = player.X + lvl.X + vx - lvl.maxX
		lvl.X = lvl.maxX
	default:
		lvl.X -= vx
	}
	switch {
	case lvl.Y-vy > 0 || player.Y < screenHalfH:
		// level coordinates should never be positives
		player.Y = player.Y - lvl.Y + vy
		lvl.Y = 0
	case lvl.Y-vy < -lvl.maxY:
		// should never be lower than the level maxs
		player.Y = player.Y + lvl.Y + vy - lvl.maxY
		lvl.Y = lvl.maxY
	default:
		lvl.Y -= vy
	}
}
