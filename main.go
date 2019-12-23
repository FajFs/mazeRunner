package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	tps          = 30
	screenWidth  = 600
	screenHeight = 600
	blockSize    = 50
	rows         = screenHeight / blockSize
	cols         = screenWidth / blockSize
)

var (
	visited = 0
)

func init() {
	rand.Seed(time.Now().UnixNano())
	m = maze{}
	stack = maze{}

	//Prepare dfsRb
	m.makeMaze(screenWidth, screenHeight)
	startIndex := rand.Intn(len(m.cells))
	m.cells[startIndex].visited = true
	stack.push(m.cells[startIndex])
	visited++
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	if visited < len(m.cells) {
		dfsRb(&m, &stack)
	} else if m.cells[0].visited {
		for i := range m.cells {
			m.cells[i].visited = false
		}
	}
	m.drawMaze(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenWidth, 1, "Maze"); err != nil {
		log.Fatal(err)
	}
}
