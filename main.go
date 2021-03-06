// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/tbruyelle/fsm"

	_ "image/png"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/sprite"
	"golang.org/x/mobile/sprite/clock"
	"golang.org/x/mobile/sprite/glsprite"
)

type Objs []*fsm.Object

func (a Objs) Remove(i int) Objs {
	a[i], a[len(a)-1], a = a[len(a)-1], nil, a[:len(a)-1]
	return a
}

var (
	startTime = time.Now()
	lastClock = clock.Time(-1)

	eng                      = glsprite.Engine()
	scene                    *fsm.Object
	tiles                    map[rune]sprite.SubTex
	objs                     Objs
	lvl                      *level
	player                   *fsm.Object
	screenW, screenH         float32
	screenHalfW, screenHalfH float32
)

func main() {
	app.Run(app.Callbacks{
		//Start: start,
		Draw:  draw,
		Touch: touch,
	})
}

func draw() {
	// Keep until golang.org/x/mogile/x11.go handle Start callback
	if scene == nil {
		start()
	}
	now := clock.Time(time.Since(startTime) * 60 / time.Second)
	if now == lastClock {
		// TODO: figure out how to limit draw callbacks to 60Hz instead of
		// burning the CPU as fast as possible.
		// TODO: (relatedly??) sync to vblank?
		return
	}
	lastClock = now

	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// test collisions

	eng.Render(scene.Node, now)
	debug.DrawFPS()
}

func touch(t event.Touch) {
	if t.Type == event.TouchEnd {
		lvl.Time = 0
		lvl.Action = &moveTo{x: float32(t.Loc.X), y: float32(t.Loc.Y)}
	}
}

func start() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	tiles = loadTiles()
	scene = &fsm.Object{Width: 1, Height: 1}
	scene.Register(nil, eng)

	// Screen dimensions
	screenW, screenH = float32(geom.Width), float32(geom.Height)
	screenHalfW, screenHalfH = screenW/2, screenH/2

	// Background
	bg := &fsm.Object{
		Width:  screenW,
		Height: screenH,
	}
	log.Println("Window", bg.Width, bg.Height)
	bg.Register(scene, eng)
	bg.Sprite = fsm.SubTex(fsm.MustLoadTexture(eng, "bg0.png"), 0, 0, 1920, 1080)

	objs = make(Objs, 0)

	// level
	lvl = loadLevel(1)
}

func loadTiles() map[rune]sprite.SubTex {
	t := fsm.MustLoadTexture(eng, "tiles.png")
	return map[rune]sprite.SubTex{
		tileFloor: fsm.SubTex(t, 32, 32, 64, 64),
		tilePlain: fsm.SubTex(t, 32, 64, 64, 96),
		tileBall:  fsm.SubTex(t, 96, 0, 128, 32),
	}
}
