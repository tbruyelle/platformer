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
		// Fix x,y to match the player dimension
		a.x, a.y = a.x-player.Width/2, a.y-player.Height/2
		// Compute vector to direction
		a.v = fsm.NewVector(a.x, a.y, player.X, player.Y)
		// Compute goal for top-left level
		a.x, a.y = lvl.X+a.v.X, lvl.Y+a.v.Y
		// Normalize vector
		a.v.Normalize()
		return
	}
	// Check if the goal is reached
	l := fsm.NewVector(a.x, a.y, o.X, o.Y).Length()
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
		log.Println("min")
		// level coordinates should never be positives
		player.X = player.X - o.X - vx
		o.X = 0
	case o.X+vx < -lvl.maxX-screenW:
		log.Println("maxX")
		// should never be lower than the level maxs
		player.X = player.X + o.X - vx - lvl.maxX - screenW
		o.X = lvl.maxX - screenW
	default:
		log.Println("default")
		o.X += vx
	}
	o.Y += vy
}
