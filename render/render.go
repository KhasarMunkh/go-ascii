package render

import (
	"github.com/KhasarMunkh/go-ascii/image_reader"
)

type Renderer interface {
	Render(p *image_reader.Pixels) (string, error) // Render the image to ASCII art and return it as a string
}

type AsciiRenderer struct {
	bw int
	bh int
}

type AnsiRenderer struct {
	bw int
	bh int
}

type BrailleRenderer struct {
	bw int
	bh int
}

func NewAsciiRenderer() *AsciiRenderer {
	return &AsciiRenderer{bw: 1, bh: 2}
}

func NewAnsiRenderer() *AnsiRenderer {
	return &AnsiRenderer{bw: 1, bh: 2}
}

func NewBrailleRenderer() *BrailleRenderer {
	return &BrailleRenderer{bw: 2, bh: 4}
}

func (r *AsciiRenderer) Render(p *image_reader.Pixels) (string, error) {
	return "ASCII Art Placeholder", nil
}

func (r *AnsiRenderer) Render(p *image_reader.Pixels) (string, error) {
	return "ANSI Art Placeholder", nil
}

func (r *BrailleRenderer) Render(p *image_reader.Pixels) (string, error) {
	return "Braille Art Placeholder", nil
}

