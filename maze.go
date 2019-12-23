package main

import (
	"github.com/hajimehoshi/ebiten"
)

var m maze
var stack maze

type maze struct {
	cells []cell
}

func (s *maze) push(t cell) {
	s.cells = append(s.cells, t)
}

func (s *maze) pop() cell {
	cell := s.cells[len(s.cells)-1]
	s.cells = s.cells[0 : len(s.cells)-1]
	return cell
}

func (s *maze) getNonVisitedNeighbours(c cell) []cell {
	var neigh []cell
	var nonVisited []cell
	index := func(i, j float64) int {
		if i < 0 || j < 0 || i > cols-1 || j > rows-1 {
			return -1
		}
		return int(i*rows + j)
	}
	if i := index(c.row, c.col-1); i != -1 {
		neigh = append(neigh, s.cells[i])
	}
	if i := index(c.row+1, c.col); i != -1 {
		neigh = append(neigh, s.cells[i])
	}
	if i := index(c.row, c.col+1); i != -1 {
		neigh = append(neigh, s.cells[i])
	}
	if i := index(c.row-1, c.col); i != -1 {
		neigh = append(neigh, s.cells[i])
	}
	for _, n := range neigh {
		if n.visited == false {
			nonVisited = append(nonVisited, n)
		}
	}
	return nonVisited
}

func remWall(a, b *cell) {
	x := a.row - b.row
	if x == 1 {
		a.walls[up] = false
		b.walls[down] = false
	} else if x == -1 {
		a.walls[down] = false
		b.walls[up] = false
	}
	y := a.col - b.col
	if y == 1 {
		a.walls[left] = false
		b.walls[right] = false
	} else if y == -1 {
		a.walls[right] = false
		b.walls[left] = false
	}
}

func (s *maze) drawMaze(screen *ebiten.Image) {
	for _, c := range s.cells {
		c.show(screen)
	}
}

func (s *maze) makeMaze(w, h float64) {
	for i := 0.0; i < rows; i++ {
		for j := 0.0; j < cols; j++ {
			s.push(newCell(i, j))
		}
	}
}
