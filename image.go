package main

import (
	"bytes"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadImage(filePath string) *ebiten.Image {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicf("Failed to load image (%s):\n%v", filePath, err)
	}

	//TODO: read non-png images
	reader := bytes.NewReader(file)
	image, err := png.Decode(reader)
	if err != nil {
		log.Panicf("Failed to decode image (%s):\n%v", filePath, err)
	}

	return ebiten.NewImageFromImage(image)
}
