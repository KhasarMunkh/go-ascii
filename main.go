package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/KhasarMunkh/go-ascii/image_reader"
	"github.com/KhasarMunkh/go-ascii/render"
	"golang.org/x/image/draw"
)

const (
	path = "assets/alex_no_background.png"
	outputFile = "output.txt"
)

func main() {
	width := flag.Int("width", 200, "Width of the output in characters")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Could not decode file", err)
	}

	// target height: keep aspect then adjust for half-block (≈1.0)
	scale := float64(*width) / float64(img.Bounds().Dx())
	tgtH := int(float64(img.Bounds().Dy()) * scale)
	dst := image.NewRGBA(image.Rect(0, 0, *width, tgtH))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	pixels, err := image_reader.DecodeImage(dst)
	if err != nil {
		log.Fatal("Could not decode image", err)
	}

	r := render.NewBrailleRenderer()
	out, err := r.Render(pixels)
	if err != nil {
		log.Fatal("Could not render image", err)
	}
	os.Stdout.WriteString(out)

	// Write the output to a file
	fout, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Could not create output file", err)
	}
	defer fout.Close()
	_, err = fout.WriteString(out)
	if err != nil {
		log.Fatal("Could not write to output file", err)
	}
}
