package main

import "math/rand"

func dfsRb(maze, stack *maze) {
	c := stack.pop()
	nonVisited := maze.getNonVisitedNeighbours(c)
	if len(nonVisited) > 0 {
		stack.push(c)
		newCell := nonVisited[rand.Intn(len(nonVisited))]
		remWall(&maze.cells[int(c.row*rows+c.col)], &maze.cells[int(newCell.row*rows+newCell.col)])
		maze.cells[int(newCell.row*rows+newCell.col)].visited = true
		stack.push(maze.cells[int(newCell.row*rows+newCell.col)])
		visited++
	}
}
