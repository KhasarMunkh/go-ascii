package image_reader

import (
	"image"
	"image/draw"
)

// Normalize converts an image to RGBA format
// and stores into a pixels slice

type Pixel struct {
	R   uint8
	G   uint8
	B   uint8
	Lum uint8
}

type Pixels struct {
	Width  int
	Height int
	Data   []Pixel
}

func NewPixels(width, height int) *Pixels {
	return &Pixels{
		Width:  width,
		Height: height,
		Data:   make([]Pixel, width*height),
	}
}

func (p *Pixels) At(x, y int) Pixel {
	return p.Data[y*p.Width+x]
}

func DecodeImage(img image.Image) (*Pixels, error) {
	rgb := normalizeToRGBA(img)
	b := rgb.Bounds()
	w := b.Max.X - b.Min.X
	h := b.Max.Y - b.Min.Y

	p := NewPixels(w, h)

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pix := rgb.RGBAAt(x, y)
			r := pix.R
			g := pix.G
			b := pix.B
			l := computeLuminance(r, g, b)

			i := y*w + x
			p.Data[i] = Pixel{
				R: r,
				G: g,
				B: b,
				Lum: l,
			}
		}
	}
	return p, nil
}

func computeLuminance(r, g, b uint8) uint8 {
	return uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func normalizeToRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	dst := image.NewRGBA(b)
	draw.Draw(dst, b, src, b.Min, draw.Src)
	return dst
}
