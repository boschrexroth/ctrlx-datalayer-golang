package plot

// Canvas describes interface for drawing graphics.
type Canvas interface {
	Bounds() Rect
	Layer(index int) Canvas
	Clip(r Rect) Canvas
	Context(r Rect) Canvas
	Text(text string, at Point, style *Style)
	Poly(points []Point, style *Style)
	Rect(r Rect, style *Style)
}
