package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"

	"github.com/KhasarMunkh/go-ascii/image_reader"
	"github.com/KhasarMunkh/go-ascii/render"
	"golang.org/x/image/draw"
)

type Mode string

const (
	ModeASCII   Mode = "ascii"
	ModeAnsi    Mode = "ansi"
	ModeBraille Mode = "braille"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: go-ascii --mode ascii|ansi|braille --width <cols> [--out <file or ->] <image file or ->\n")
	os.Exit(2)
}

func main() {
	mode := flag.String("mode", string(ModeBraille), "ascii|ansi|braille")
	width := flag.Int("width", 200, "width in character cells")
	outpath := flag.String("out", "", "output file (default stdout). Use - for stdout.")
	flag.Parse()

	// resolve positional input (path or "-")
	inputPath := "-"
	switch flag.NArg() {
	case 0: // stdin
	case 1:
		inputPath = flag.Arg(0)
	default:
		usage()
	}

	// validate flags
	m := Mode(*mode)
	if m != ModeASCII && m != ModeAnsi && m != ModeBraille {
		log.Fatalf("invalid --mode: %q (use ascii|ansi|braille)", *mode)
	}
	if *width <= 0 {
		log.Fatalf("--width must be > 0")
	}

	// load image (file or stdin)
	img, err := loadImage(inputPath)
	if err != nil {
		log.Fatalf("could not read image: %v", err)
	}

	// cell geometry per mode
	cellW, cellH := cellGeometryFor(m)

	// target pixel size from character grid
	cols := *width
	dstW := max(1, cols*cellW)

	srcW := img.Bounds().Dx()
	srcH := img.Bounds().Dy()
	if srcW == 0 || srcH == 0 {
		log.Fatal("empty image bounds")
	}

	// preserve aspect: dstH by proportional scaling, then snap to multiple of cellH
	dstH := int(float64(dstW) * float64(srcH) / float64(srcW))
	if dstH < cellH {
		dstH = cellH
	}
	dstH = (dstH / cellH) * cellH
	if dstH == 0 {
		dstH = cellH
	}

	// resize to target pixels
	dst := image.NewRGBA(image.Rect(0, 0, dstW, dstH))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	// convert to your Pixels type
	pixels, err := image_reader.DecodeImage(dst)
	if err != nil {
		log.Fatalf("could not decode image to pixels: %v", err)
	}

	// render
	var out string
	switch m {
	case ModeBraille:
		out, err = render.NewBrailleRenderer().Render(pixels) // uses 2x4 cells internally
	case ModeASCII:
		out, err = render.NewAsciiRenderer().Render(pixels) // uses 1x2 cells internally
	case ModeAnsi:
		out, err = render.NewAnsiRenderer().Render(pixels) // uses 1x2 cells internally
	}
	if err != nil {
		log.Fatalf("render error: %v", err)
	}

	// write output (stdout if empty or "-")
	writeOut(*outpath, out)
}

// ----- helpers -----

func loadImage(path string) (image.Image, error) {
	var r io.Reader
	if path == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r = f
	}
	img, _, err := image.Decode(r)
	return img, err
}

// ASCII/ANSI ~ 1x2 cells; Braille 2x4 dots per char
func cellGeometryFor(m Mode) (cellW, cellH int) {
	switch m {
	case ModeBraille:
		return 2, 4
	default: // ascii, ansi
		return 1, 2
	}
}

func writeOut(outpath, output string) {
	if outpath == "" || outpath == "-" {
		_, _ = os.Stdout.WriteString(output)
		return
	}
	fout, err := os.Create(outpath)
	if err != nil {
		log.Fatalf("could not create output file: %v", err)
	}
	defer fout.Close()
	if _, err := fout.WriteString(output); err != nil {
		log.Fatalf("could not write to output file: %v", err)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
