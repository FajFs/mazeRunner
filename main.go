package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const (
	screenWidth  = 600
	screenHeight = 600
	blockSize    = 50
	rows         = screenHeight / blockSize
	cols         = screenWidth / blockSize
)

var (
	visited    = 0
	normalFont font.Face
)

func init() {
	ebiten.SetMaxTPS(30)
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	normalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	rand.Seed(time.Now().UnixNano())
	m = maze{}
	s = maze{}

	m.makeMaze(screenWidth, screenHeight)
	startIndex := rand.Intn(len(m.cells))
	visited++
	m.cells[startIndex].visited = true
	s.push(m.cells[startIndex])
}

func update(screen *ebiten.Image) error {
	fps := ebiten.CurrentTPS()
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	text.Draw(screen, fmt.Sprintf("FPS: %0.2f", fps), normalFont, 20, 40, color.White)
	if visited < len(m.cells) {
		updateMaze(&m, &s)
	} else {
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
