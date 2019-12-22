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
	screenWidth  = 1000
	screenHeight = 1000
)

var (
	normalFont font.Face
)

func init() {
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
	makeMaze(screenWidth, screenHeight)
	rand.Seed(time.Now().UnixNano())
	maze[rand.Intn(len(maze))].visited = true
	// currCell = maze[rand.Intn(len(maze))]
}

func update(screen *ebiten.Image) error {
	fps := ebiten.CurrentTPS()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	text.Draw(screen, fmt.Sprintf("FPS: %0.2f", fps), normalFont, 20, 40, color.White)
	drawMaze(screen)
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenWidth, 1, "Maze"); err != nil {
		log.Fatal(err)
	}
}
