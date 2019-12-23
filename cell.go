package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
	x := c.row * blockSize
	y := c.col * blockSize
	w := float64(blockSize)
	if c.walls[0] {
		ebitenutil.DrawLine(s, x, y, x+w, y, color.White)
	}
	if c.walls[1] {
		ebitenutil.DrawLine(s, x+w, y, x+w, y+w, color.White)
	}
	if c.walls[2] {
		ebitenutil.DrawLine(s, x+w, y+w, x, y+w, color.White)
	}
	if c.walls[3] {
		ebitenutil.DrawLine(s, x, y+w, x, y, color.White)
	}
	if c.visited {
		ebitenutil.DrawRect(s, c.col*blockSize+10, c.row*blockSize+10, blockSize-20, blockSize-20, color.White)
	}
}
