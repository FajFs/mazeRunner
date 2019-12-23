package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

var m maze
var s maze

type maze struct {
	cells []cell
}

// Push adds an Item to the top of the stack
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
		a.walls[3] = false
		b.walls[1] = false
	} else if x == -1 {
		a.walls[1] = false
		b.walls[3] = false
	}
	y := a.col - b.col
	if y == 1 {
		a.walls[0] = false
		b.walls[2] = false
	} else if y == -1 {
		a.walls[2] = false
		b.walls[0] = false
	}
}

func updateMaze(maze, stack *maze) {
	/*
		pop a cell from the stack and make it a current cell
		If the current cell has any neighbours which have not been visited
		Push the current cell to the stack
		Choose one of the unvisited neighbours
		Remove the wall between the current cell and the chosen cell
		Mark the chosen cell as visited and push it to the stack
	*/
	c := stack.pop()
	nonVisited := maze.getNonVisitedNeighbours(c)
	if len(nonVisited) > 0 {
		stack.push(c)
		i := rand.Intn(len(nonVisited))
		nC := nonVisited[i]

		remWall(&maze.cells[int(c.row+c.col*cols)], &maze.cells[int(nC.row+nC.col*cols)])
		maze.cells[int(nC.row*rows+nC.col)].visited = true
		visited++
		stack.push(maze.cells[int(nC.row*rows+nC.col)])
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
