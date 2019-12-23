package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	left = iota
	down
	right
	up
)

type cell struct {
	visited  bool
	walls    [4]bool
	row, col float64
}

func newCell(row, col float64) cell {
	return cell{row: row, col: col, visited: false, walls: [4]bool{true, true, true, true}}
}

func (c cell) show(s *ebiten.Image) {
	x := c.col * blockSize
	y := c.row * blockSize
	w := float64(blockSize)
	if c.walls[up] {
		ebitenutil.DrawLine(s, x, y, x+w, y, color.White)
	}
	if c.walls[right] {
		ebitenutil.DrawLine(s, x+w, y, x+w, y+w, color.White)
	}
	if c.walls[down] {
		ebitenutil.DrawLine(s, x+w, y+w, x, y+w, color.White)
	}
	if c.walls[left] {
		ebitenutil.DrawLine(s, x, y+w, x, y, color.White)
	}
	if c.visited {
		ebitenutil.DrawRect(s, c.col*blockSize+4, c.row*blockSize+4, blockSize-8, blockSize-8, color.White)
	}
}
