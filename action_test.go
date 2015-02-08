package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tbruyelle/fsm"
	"testing"
)

func TestScrollLevelForward(t *testing.T) {
	lvl = &level{minX: -300, minY: -300}
	screenW = 200
	screenH = 200
	player = &fsm.Object{X: 100, Y: 100}

	scroll(10, 10)

	assert.Equal(t, -10, lvl.X)
	assert.Equal(t, -10, lvl.Y)
	assert.Equal(t, 100, player.X)
	assert.Equal(t, 100, player.Y)
}

func TestScrollLevelBackward(t *testing.T) {
	lvl = &level{minX: -300, minY: -300}
	lvl.X, lvl.Y = -20, -20
	screenW = 200
	screenH = 200
	player = &fsm.Object{X: 100, Y: 100}

	scroll(-10, -10)

	assert.Equal(t, -10, lvl.X)
	assert.Equal(t, -10, lvl.Y)
	assert.Equal(t, 100, player.X)
	assert.Equal(t, 100, player.Y)
}

func TestScrollPlayerForward(t *testing.T) {
	lvl = &level{minX: -300, minY: -300}
	lvl.X, lvl.Y = -300, -300
	screenW = 200
	screenH = 200
	player = &fsm.Object{X: 100, Y: 100}

	scroll(10, 10)

	assert.Equal(t, -300, lvl.X)
	assert.Equal(t, -300, lvl.Y)
	assert.Equal(t, 110, player.X)
	assert.Equal(t, 110, player.Y)
}

func TestScrollPlayerBackward(t *testing.T) {
	lvl = &level{minX: -300, minY: -300}
	screenW = 200
	screenH = 200
	player = &fsm.Object{X: 100, Y: 100}

	scroll(-10, -10)

	assert.Equal(t, 0, lvl.X)
	assert.Equal(t, 0, lvl.Y)
	assert.Equal(t, 90, player.X)
	assert.Equal(t, 90, player.Y)
}

func TestScrollLevelPlayerForward(t *testing.T) {
	lvl = &level{minX: -300, minY: -300}
	lvl.X, lvl.Y = -295, -295
	screenW = 200
	screenH = 200
	player = &fsm.Object{X: 100, Y: 100}

	scroll(10, 10)

	assert.Equal(t, -300, lvl.X)
	assert.Equal(t, -300, lvl.Y)
	assert.Equal(t, 105, player.X)
	assert.Equal(t, 105, player.Y)
}

func TestScrollLevelPlayerBackward(t *testing.T) {
	lvl = &level{minX: -300, minY: -300}
	lvl.X, lvl.Y = -5, -5
	screenW = 200
	screenH = 200
	player = &fsm.Object{X: 100, Y: 100}

	scroll(-10, -10)

	assert.Equal(t, 0, lvl.X)
	assert.Equal(t, 0, lvl.Y)
	assert.Equal(t, 95, player.X)
	assert.Equal(t, 95, player.Y)
}
