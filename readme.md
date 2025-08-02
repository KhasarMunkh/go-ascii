1) # Go Image Processing Resources 
# Go Image Processing
https://pkg.go.dev/image

# go blog on image processing
https://go.dev/blog/image
https://go.dev/blog/image-draw
2) # Tone mapping & dithering (to avoid banding)
# Floyd Steinberg dithering
https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering
# Ordered dithering
https://en.wikipedia.org/wiki/Ordered_dithering
# Atkinson dithering
https://en.wikipedia.org/wiki/Atkinson_dithering

3) # Terminals, color, and glyph sets
# ANSI escape codes (SGR) for color
https://en.wikipedia.org/wiki/ANSI_escape_code
# Unicode Block Elements - courser pixels than ascii
https://en.wikipedia.org/wiki/Block_Elements
# Unicode Braille Patterns - finer pixels than ascii
https://en.wikipedia.org/wiki/Braille_Patterns

4) # Character ramps and density
https://paulbourke.net/dataformats/asciiart

TODO:
A 3-step “first build” using the docs above

    Decode + resize with x/image/draw (CatmullRom); keep width fixed and adjust height by an aspect ratio (e.g., ~0.5) to match terminal cells.
    Go Packages

    Compute luminance from image/color RGBA (remember 16-bit → 8-bit) and map to a Bourke ramp.
    Go Packages
    paulbourke.net

    Emit text (plain, ANSI 24-bit, or both). For color, use SGR 38;2;R;G;B and reset after each glyph or run.
    Wikipedia
