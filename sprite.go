package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func getSprites(spriteLocation string, width int, height int, spriteState SpriteState) (Sprite, error) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	spriteImage, err := os.Open(spriteLocation)
	if err != nil {
		log.Println("No file?")
	}
	defer spriteImage.Close()

	pixels, err := ConvertPixels(spriteImage)
	if err != nil {
		log.Println("Can't convert?", width, height)
	}

	totalSprites := len(pixels[0]) / width
	sprites := make([][][]Pixel, totalSprites)

	for i := 0; i < totalSprites; i++ {
		sprites[i] = make([][]Pixel, width)
		for y := 0; y < len(pixels); y++ {
			sprites[i][y] = make([]Pixel, width)
			for x := 0; x < width; x++ {
				sprites[i][y][x] = pixels[y][x]

			}
			pixels[y] = append(pixels[y][width:])
		}
	}

	var Sprite Sprite
	Sprite.sprites = sprites
	if len(sprites) == 0 {
		Sprite.animated = false
	} else {
		Sprite.animated = true
	}
	Sprite.spriteState = spriteState

	return Sprite, nil

}

func ConvertPixels(fileImage io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(fileImage)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	// reverse array cause upsidedown
	for i, j := 0, len(pixels)-1; i < j; i, j = i+1, j-1 {
		pixels[i], pixels[j] = pixels[j], pixels[i]
	}

	return pixels, nil
}
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

type Pixel struct {
	R int
	G int
	B int
	A int // if 0 - invis
}

type Sprite struct {
	sprites     [][][]Pixel
	spriteState SpriteState
	animated    bool
}
type SpriteState struct {
	idle   []int
	moving []int
	custom []int
}
