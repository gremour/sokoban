package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/gremour/sokoban/game"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sokoban",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pic, err := loadPicture("sprites.png")
	if err != nil {
		panic(err)
	}

	var sprites []*pixel.Sprite
	for i := 0; i < int(pic.Bounds().Max.X/32); i++ {
		sprite := pixel.NewSprite(pic, pixel.R(float64(i*32), 0, float64((i+1)*32), 32))
		sprites = append(sprites, sprite)
	}

	g, err := game.New("level", sprites)
	if err != nil {
		log.Fatal(err)
	}

	wmap := float64(g.Map.Width) * 32
	hmap := float64(g.Map.Height) * 32
	dx := (win.Bounds().W() - wmap) / 2
	dy := (win.Bounds().H() - hmap) / 2

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	winText := text.New(win.Bounds().Center().Add(pixel.V(-175, 0)), atlas)
	winText.Color = colornames.Lightsalmon
	fmt.Fprintf(winText, "You Win!")

	matr := pixel.IM.Moved(pixel.V(dx, dy))
	for !win.Closed() {
		win.Clear(colornames.Darkgray)
		g.Draw(win, matr)
		g.DrawPlayer(win, matr)

		if g.Win {
			winText.Draw(win, pixel.IM.Scaled(winText.Orig, 6))
		} else {
			switch {
			case win.JustPressed(pixelgl.KeyLeft):
				g.Move(-1, 0)
			case win.JustPressed(pixelgl.KeyRight):
				g.Move(1, 0)
			case win.JustPressed(pixelgl.KeyDown):
				g.Move(0, 1)
			case win.JustPressed(pixelgl.KeyUp):
				g.Move(0, -1)
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
