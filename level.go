package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/mobile/app"
)

type level struct {
	num   int
	tiles [][]rune
}

func loadLevel(num int) *level {
	l := &level{num: num}
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
		}
	}
	return l
}
