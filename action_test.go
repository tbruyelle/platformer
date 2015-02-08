package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/tbruyelle/fsm"
	"testing"
)

func init() {
	lvl = &level{minX: -300, minY: -300}
	screenW = 200
	screenH = 200
	screenHalfW = screenW / 2
	screenHalfH = screenH / 2
}

func TestScrollForwardFromStart(t *testing.T) {
	player = &fsm.Object{X: 100, Y: 100}

	scroll(10, 10)

	assert.Equal(t, -10, lvl.X)
	assert.Equal(t, -10, lvl.Y)
	assert.Equal(t, 100, player.X)
	assert.Equal(t, 100, player.Y)
}

func TestScrollBackwardFrom20(t *testing.T) {
	lvl.X, lvl.Y = -20, -20
	player = &fsm.Object{X: 100, Y: 100}

	scroll(-10, -10)

	assert.Equal(t, -10, lvl.X)
	assert.Equal(t, -10, lvl.Y)
	assert.Equal(t, 100, player.X)
	assert.Equal(t, 100, player.Y)
}

func TestScrollForwardFromEnd(t *testing.T) {
	lvl.X, lvl.Y = -300, -300
	player = &fsm.Object{X: 100, Y: 100}

	scroll(10, 10)

	assert.Equal(t, -300, lvl.X)
	assert.Equal(t, -300, lvl.Y)
	assert.Equal(t, 110, player.X)
	assert.Equal(t, 110, player.Y)
}

func TestScrollBackwardFromStart(t *testing.T) {
	player = &fsm.Object{X: 100, Y: 100}

	scroll(-10, -10)

	assert.Equal(t, 0, lvl.X)
	assert.Equal(t, 0, lvl.Y)
	assert.Equal(t, 90, player.X)
	assert.Equal(t, 90, player.Y)
}

func TestScrollForwardFromNearEnd(t *testing.T) {
	lvl.X, lvl.Y = -295, -295
	player = &fsm.Object{X: 100, Y: 100}

	scroll(10, 10)

	assert.Equal(t, -300, lvl.X)
	assert.Equal(t, -300, lvl.Y)
	assert.Equal(t, 105, player.X)
	assert.Equal(t, 105, player.Y)
}

func TestScrollBackwardFromNearStart(t *testing.T) {
	lvl.X, lvl.Y = -5, -5
	player = &fsm.Object{X: 100, Y: 100}

	scroll(-10, -10)

	assert.Equal(t, 0, lvl.X)
	assert.Equal(t, 0, lvl.Y)
	assert.Equal(t, 95, player.X)
	assert.Equal(t, 95, player.Y)
}

func TestScrollForwardFromNearStart(t *testing.T) {
	player = &fsm.Object{X: 95, Y: 95}

	scroll(10, 10)

	assert.Equal(t, -5, lvl.X)
	assert.Equal(t, -5, lvl.Y)
	assert.Equal(t, 100, player.X)
	assert.Equal(t, 100, player.Y)
}

func TestScrollBackwardFromNearEnd(t *testing.T) {
	lvl.X, lvl.Y = -300, -300
	player = &fsm.Object{X: 105, Y: 105}

	scroll(-10, -10)

	assert.Equal(t, -295, lvl.X)
	assert.Equal(t, -295, lvl.Y)
	assert.Equal(t, 100, player.X)
	assert.Equal(t, 100, player.Y)
}
