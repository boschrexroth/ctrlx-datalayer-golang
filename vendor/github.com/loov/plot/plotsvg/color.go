package plotsvg

import (
	"fmt"
	"image/color"
)

// convertColorToHex converts color to an hex encoded string.
func convertColorToHex(color color.Color) string {
	r, g, b, a := color.RGBA()
	if a > 0 {
		r, g, b, a = r*0xff/a, g*0xff/a, b*0xff/a, a/0xff
		if r > 0xFF {
			r = 0xFF
		}
		if g > 0xFF {
			g = 0xFF
		}
		if b > 0xFF {
			b = 0xFF
		}
		if a > 0xFF {
			a = 0xFF
		}
	} else {
		r, g, b, a = 0, 0, 0, 0
	}
	hex := r<<24 | g<<16 | b<<8 | a
	return fmt.Sprintf("#%08x", hex)
}
