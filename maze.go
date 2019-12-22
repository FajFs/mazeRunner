package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var maze []cell
var currCell cell

type cell struct {
	visited bool
	walls   [4]bool
	i, j    float64
}

const (
	blockSize = 100
)

func newBlock(i, j float64) cell {
	return cell{i: i, j: j, visited: false, walls: [4]bool{true, true, true, true}}
}

func (c cell) show(s *ebiten.Image) {
	x := c.i * blockSize
	y := c.j * blockSize
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
		ebitenutil.DrawRect(s, c.j*blockSize+10, c.i*blockSize+10, blockSize-20, blockSize-20, color.White)
	}
}

func drawMaze(screen *ebiten.Image) {
	for _, c := range maze {
		c.show(screen)
	}
}

func makeMaze(w, h float64) {
	for i := 0.0; i < h/blockSize; i++ {
		for j := 0.0; j < w/blockSize; j++ {
			maze = append(maze, newBlock(i, j))
		}
	}
}
