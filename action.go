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
		a.v = fsm.NewVector(a.x, a.y, player.X, player.Y)
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
	// Where to apply those values ?
	switch {
	case o.X+vx > 0:
		// level coordinates should never be positives
		player.X = player.X - o.X - vx
		o.X = 0
	case o.X+vx < -lvl.maxX-screenW:
		// should never be lower than the level maxs
		player.X = player.X + o.X - vx - lvl.maxX - screenW
		o.X = lvl.maxX - screenW
	default:
		o.X += vx
	}
	switch {
	case o.Y+vy > 0:
		// level coordinates should never be positives
		player.Y = player.Y - o.Y - vy
		o.Y = 0
	case o.Y+vy < -lvl.maxY-screenH:
		// should never be lower than the level maxs
		player.Y = player.Y + o.Y - vy - lvl.maxY - screenH
		o.Y = lvl.maxY - screenH
	default:
		o.Y += vy
	}
}
