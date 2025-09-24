package identicon

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/fogleman/gg"
)

func New(data []byte, size int) image.Image {
	var w, h int = size, size
	block := float64(size) / float64(5)
	fmt.Printf("block size: %f\n", block)

	hash := md5.Sum(data)
	// Use first 3 bytes of hash for RGB color
	fillColor := color.RGBA{hash[0], hash[1], hash[2], 255}

	dc := gg.NewContext(w, h)
	dc.SetColor(color.White)
	dc.Clear()
	dc.SetColor(fillColor)

	var x, y float64
	for y = 0; y < 5; y++ {
		for x = 0; x < 3; x++ {
			i := int(x + y*3)
			if hash[i]%2 == 0 {
				dc.DrawRectangle(x*block, y*block, block, block)
				dc.Fill()
				// mirror for symmetry
				dc.DrawRectangle((4-x)*block, y*block, block, block)
				dc.Fill()
			}
		}
	}

	return dc.Image()
}

func SavePNG(img image.Image, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		return err
	}
	return nil
}
