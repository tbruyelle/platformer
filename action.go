package main

import (
	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/sprite/clock"
	"math"
)

type moveTo struct {
	x, y   float32
	vx, vy float32
}

const (
	playerSpeed = 5
)

func (a *moveTo) Do(o *fsm.Object, t clock.Time) {
	if o.Time == 0 {
		o.Time = t
		// Compute vector to direction
		a.vx = a.x - o.Width/2 - o.X
		a.vy = a.y - o.Width/2 - o.Y
		// Compute vector length
		l := float32(math.Sqrt(float64(a.vx*a.vx + a.vy*a.vy)))
		// Normalize vector
		a.vx, a.vy = a.vx/l, a.vy/l
	}
	f := clock.EaseIn(o.Time, o.Time+10, t)
	o.Vx, o.Vy = a.vx*playerSpeed*f, a.vy*playerSpeed*f
	if o.X == a.x && o.Y == a.y {
		// move over
		o.Reset()
		o.Action = nil
	}

}
