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
	//log.Println("player", player.X, player.Y)
	//log.Println("v", v.X, a.v.Y, l)
	//log.Println("-------")
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
	dx, dy := lvl.X-vx, lvl.Y-vy
	log.Println("dx,dy", dx, dy)
	pdx, pdy := screenHalfW-player.X, screenHalfH-player.Y
	log.Println("pdx,pdy", pdx, pdy)
	switch {
	case vx < 0:
		// scroll left
		switch {
		case dx > 0:
			/// left limit is reached
			lvl.X = 0
			player.X -= dx
		case pdx != 0:
			// player is not centered
			if pdx < vx {
				player.X += vx
			} else {
				player.X = screenHalfW
				lvl.X += (pdx - vx)
			}
		default:
			lvl.X = dx
		}
	case vx > 0:
		// scroll right
		switch {
		case dx < lvl.minX:
			// right limit is reached
			lvl.X = lvl.minX
			player.X += -dx + lvl.minX
		case pdx != 0:
			// player is not centered
			if pdx > vx {
				player.X += vx
			} else {
				player.X = screenHalfW
				lvl.X -= (pdx - vx)
			}
		default:
			lvl.X = dx
		}
	}
	switch {
	case vy < 0:
		// scroll up
		switch {
		case dy > 0:
			// UP limit is reached
			lvl.Y = 0
			player.Y -= dy
		case pdy != 0:
			// player is not centered
			if pdy < vy {
				player.Y += vy
			} else {
				player.Y = screenHalfH
				lvl.Y += (pdy - vy)
			}
		default:
			lvl.Y = dy
		}
	case vy > 0:
		// scroll down
		switch {
		case dy < lvl.minY:
			// down limit is reached
			lvl.Y = lvl.minY
			player.Y += -dy + lvl.minY
		case pdy != 0:
			// player is not centered
			if pdy > vy {
				player.Y += vy
			} else {
				player.Y = screenHalfH
				lvl.Y -= (pdy - vy)
			}
		default:
			lvl.Y = dy
		}
	}
}
