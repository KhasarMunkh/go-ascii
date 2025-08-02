package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"github.com/KhasarMunkh/go-ascii/image_reader"
)

func main() {
	f, err := os.Open("assets/zoro.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Could not decode file", err)
	}

	pixels, err:= image_reader.DecodeImage(img)
	if err != nil {
		log.Fatal("Could not decode image", err)
	}


}

/*
type Image interface {
	toRGBA() *image.RGBA 
	toGray() *image.Gray 
}



type Pixel struct {
	lum  uint8
}

type Pixels []Pixel 

func main() {
	f, err := os.Open("assets/zoro.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		log.Fatal("Could not decode file", err)
	}
	fmt.Println("format: ", format)
	p := Pixels{}
	rgba := Normalize(img)
	b := rgba.Bounds()
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			p = append(p, Pixel{
				lum: calculateLum(rgba.At(x, y)),
			})
		}
	}
	fmt.Println("Pixels length: ", len(p))
	// Example of how to use the Pixels
	for i, pixel := range p {
		if i < 10 { // Print only the first 10 pixels
			fmt.Printf("Pixel %d: Luminosity %d\n", i, pixel.lum)
		}
	}
}

func calculateLum(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	// Using the luminosity method to calculate perceived brightness
	lum := uint8(0.2126*float64(r>>8) + 0.7152*float64(g>>8) + 0.0722*float64(b>>8))
	return lum
}


func Normalize(img image.Image) *image.RGBA {
	b := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	return dst
}

func PrintAscii() {

}
*/
