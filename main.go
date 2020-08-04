package main

import (
	"image"
	"log"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

	brick := pixel.NewSprite(pic, pixel.R(0, 0, 32, 32))

	center := pixel.IM.Moved(win.Bounds().Center())
	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		brick.Draw(win, center)
		brick.Draw(win, center.Moved(pixel.V(32, 0)))
		brick.Draw(win, center.Moved(pixel.V(64, 0)))

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
