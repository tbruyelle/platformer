package main

import (
	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/sprite/clock"
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
		a.x, a.y = a.x-o.Width/2, a.y-o.Height/2
		// Compute vector to direction
		a.v = fsm.NewVector(o.X, o.Y, a.x, a.y)
		// Normalize vector
		a.v.Normalize()
	}
	// Check if the goal is reached
	if fsm.NewVector(o.X, o.Y, a.x, a.y).Length() <= playerSpeed {
		// move over
		o.Reset()
		o.Action = nil
		return
	}
	f := clock.EaseIn(o.Time, o.Time+10, t)
	o.Vx, o.Vy = a.v.X*playerSpeed*f, a.v.Y*playerSpeed*f

	lvl.playerX, lvl.playerY = o.Vx+o.X, o.Vy+o.Y
	lvl.playerMove()
}
