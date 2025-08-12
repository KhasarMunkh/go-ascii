package render

import (
	"fmt"
	"strings"

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

var (
	asciiTableSimple   = " .:-=+*#%@"
	asciiTableDetailed = " .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
)

func (r *AsciiRenderer) Render(p *image_reader.Pixels) (string, error) {
	if p.Width == 0 || p.Height < 2 {
		return "", fmt.Errorf("empty or too-small image")
	}

	var sb strings.Builder
	// Rough capacity: each cell ≈ 24 bytes of escape codes + glyph,
	// plus newline per text row.
	sb.Grow(p.Width * p.Height * 12)

	for y := 0; y < p.Height-1; y += 2 { // take two source rows per text row
		for x := 0; x < p.Width; x++ {
			top := p.At(x, y)
			bot := p.At(x, y+1)

			tl := top.Lum
			b1 := bot.Lum
			al := (tl + b1) / 2                              // average luminance for the half-block
			i := int(al) * (len(asciiTableDetailed) - 1) / 255 // map to index
			sb.WriteByte(asciiTableDetailed[i])
		}
		sb.WriteByte('\n') // newline at end of text row
	}

	return sb.String(), nil
}

func (r *AsciiRenderer) RenderColor(p *image_reader.Pixels) (string, error) {
	if p.Width == 0 || p.Height < 2 {
		return "", fmt.Errorf("empty or too-small image")
	}

	var sb strings.Builder
	// Rough capacity: each cell ≈ 24 bytes of escape codes + glyph,
	// plus newline per text row.
	sb.Grow(p.Width * p.Height * 12)

	for y := 0; y < p.Height-1; y += 2 { // take two source rows per text row
		for x := 0; x < p.Width; x++ {
			top := p.At(x, y)
			bot := p.At(x, y+1)

			tl := top.Lum
			b1 := bot.Lum
			al := (tl + b1) / 2                              // average luminance for the half-block
			i := int(al) * (len(asciiTableDetailed) - 1) / 255 // map to index
			glyph := asciiTableDetailed[i]

			// FG (top), BG (bottom), then the glyph
			sb.WriteString(fg256(top.R, top.G, top.B)) // foreground
			sb.WriteByte(glyph)
		}
		sb.WriteString("\x1b[0m") // reset colors
		sb.WriteByte('\n') // newline at end of text row
	}

	return sb.String(), nil
}


func rgbToAnsi256(r, g, b uint8) uint8 {
    r6, g6, b6 := int(r/51), int(g/51), int(b/51) // 0–5 cube
    return uint8(16 + 36*r6 + 6*g6 + b6)
}
func fg256(r, g, b uint8) string {
    return fmt.Sprintf("\x1b[38;5;%dm", rgbToAnsi256(r, g, b))
}

// Render the image to ANSI art using half-block characters
// Each text row is made of two source rows, using the top pixel for foreground
// and the bottom pixel for background color.
// The half-block character "▀" is used to represent the two pixels.

func (r *AnsiRenderer) Render(p *image_reader.Pixels) (string, error) {
	if p.Width == 0 || p.Height < 2 {
		return "", fmt.Errorf("empty or too-small image")
	}

	var sb strings.Builder
	// Rough capacity: each cell ≈ 24 bytes of escape codes + glyph,
	// plus newline per text row.
	sb.Grow(p.Width * p.Height * 12)

	for y := 0; y < p.Height-1; y += 2 { // take two source rows per text row
		for x := 0; x < p.Width; x++ {
			top := p.At(x, y)
			bot := p.At(x, y+1)

			// FG (top), BG (bottom), then the half-block “▀”
			fmt.Fprintf(&sb,
				"\x1b[38;2;%d;%d;%dm"+ // foreground
					"\x1b[48;2;%d;%d;%dm"+ // background
					"▀",
				top.R, top.G, top.B,
				bot.R, bot.G, bot.B,
			)
		}
		sb.WriteString("\x1b[0m\n") // reset + newline at end of text row
	}

	return sb.String(), nil
}

func (r *BrailleRenderer) Render(p *image_reader.Pixels) (string, error) {
	if p.Width == 0 || p.Height < 4 {
		return "", fmt.Errorf("empty or too-small image")
	}
	// Braille characters are 2x4 blocks, so we need to adjust the width and height
	bw := p.Width / 2
	bh := p.Height / 4
	if bw == 0 || bh == 0 {
		return "", fmt.Errorf("image too small for braille rendering")
	}
	var sb strings.Builder
	// Rough capacity: each cell ≈ 24 bytes of escape codes + glyph,
	// plus newline per text row.
	sb.Grow(bw * bh * 12)
	for y := 0; y < p.Height-3; y += 4 { // take four source rows per text row
		for x := 0; x < p.Width-1; x += 2 {
			// Get the 2x4 block of pixels
			pix := [8]image_reader.Pixel{
				p.At(x, y),     // 0
				p.At(x+1, y),   // 1
				p.At(x, y+1),   // 2
				p.At(x+1, y+1), // 3
				p.At(x, y+2),   // 4
				p.At(x+1, y+2), // 5
				p.At(x, y+3),   // 6
				p.At(x+1, y+3), // 7
			}
			// Calculate the braille character from the pixels
			var brailleChar rune
			for i, p := range pix {
				if p.Lum > 127 { // threshold for "filled" pixel
					brailleChar |= 1 << i // set the bit for this pixel
				}
			}
			// Add the braille character to the string builder
			sb.WriteRune(0x2800 + brailleChar) // Braille characters start
			// at U+2800
		}
		sb.WriteByte('\n') // newline at end of text row
	}
	return sb.String(), nil
}
