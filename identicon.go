package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Identicon represents an identicon with its configuration
type Identicon struct {
	Hash  []byte
	Color color.RGBA
	Grid  []bool
	Image *image.RGBA
}

// New creates a new identicon from an input string
func New(input string) *Identicon {
	hash := md5.Sum([]byte(input))
	
	identicon := &Identicon{
		Hash: hash[:],
	}
	
	identicon.pickColor()
	identicon.buildGrid()
	identicon.buildImage()
	
	return identicon
}

// pickColor selects a color based on the first 3 bytes of the hash
func (i *Identicon) pickColor() {
	i.Color = color.RGBA{
		R: i.Hash[0],
		G: i.Hash[1], 
		B: i.Hash[2],
		A: 255,
	}
}

// buildGrid creates a 5x5 grid pattern based on the hash
func (i *Identicon) buildGrid() {
	grid := make([]bool, 25)
	
	// Use bytes 3-14 of the hash to determine the pattern
	// We only need the first 15 positions (3 columns) since we mirror
	for idx := 0; idx < 15; idx++ {
		if i.Hash[idx%len(i.Hash)] % 2 == 0 {
			grid[idx] = true
		}
	}
	
	// Mirror the first 3 columns to create symmetry
	for row := 0; row < 5; row++ {
		grid[row*5+3] = grid[row*5+1] // Mirror column 1 to column 3
		grid[row*5+4] = grid[row*5+0] // Mirror column 0 to column 4
	}
	
	i.Grid = grid
}

// buildImage creates the actual image from the grid
func (i *Identicon) buildImage() {
	const pixelSize = 50
	const imageSize = pixelSize * 5
	
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	
	// Fill with white background
	for y := 0; y < imageSize; y++ {
		for x := 0; x < imageSize; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	
	// Draw the pattern
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if i.Grid[row*5+col] {
				// Fill the pixel block
				for y := row * pixelSize; y < (row+1)*pixelSize; y++ {
					for x := col * pixelSize; x < (col+1)*pixelSize; x++ {
						img.Set(x, y, i.Color)
					}
				}
			}
		}
	}
	
	i.Image = img
}

// Save saves the identicon to a PNG file
func (i *Identicon) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	
	err = png.Encode(file, i.Image)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}
	
	return nil
}