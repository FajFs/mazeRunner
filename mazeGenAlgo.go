package main

import "math/rand"

/*
1.Choose the initial cell, mark it as visited and push it to the stack
2.While the stack is not empty
  1. Pop a cell from the stack and make it a current cell
  2. If the current cell has any neighbours which have not been visited
    1. Push the current cell to the stack
    2. Choose one of the unvisited neighbours
    3. Remove the wall between the current cell and the chosen cell
    4. Mark the chosen cell as visited and push it to the stack
*/

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
