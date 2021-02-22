package plot

// Plot defines a combination of elements that can be drawn to the canvas.
type Plot struct {
	// X, Y are the axis information
	X, Y   *Axis
	Margin Rect
	Elements
	// DefaultStyle
	Theme
}

// Element is a drawable plot element.
type Element interface {
	Draw(plot *Plot, canvas Canvas)
}

// Dataset represents an Element that contains data
type Dataset interface {
	Element
	// TODO: remove and replace with recommended Axis
	Stats() Stats
}

// New creates a new empty plot.
func New() *Plot {
	x, y := NewAxis(), NewAxis()
	y.Flip = true

	return &Plot{
		X:     x,
		Y:     y,
		Theme: NewTheme(),
	}
}

// Draw draws plot to the specified canvas, creating axes automatically when necessary.
func (plot *Plot) Draw(canvas Canvas) {
	if !plot.X.IsValid() || !plot.Y.IsValid() {
		tmpplot := &Plot{}
		*tmpplot = *plot
		plot = tmpplot
		plot.X, plot.Y = detectAxis(plot.X, plot.Y, plot.Elements)
	}

	bounds := canvas.Bounds()
	if !plot.Margin.Empty() {
		bounds = bounds.Inset(plot.Margin)
	}

	for _, element := range plot.Elements {
		element.Draw(plot, canvas.Context(bounds))
	}
}

// AxisGroup allows sub-elements to have different different axes defined rather than the top-level plot.
type AxisGroup struct {
	X, Y *Axis
	Elements
}

// NewAxisGroup creates a new axis group.
func NewAxisGroup(els ...Element) *AxisGroup {
	x, y := NewAxis(), NewAxis()
	y.Flip = true
	return &AxisGroup{
		X: x, Y: y,
		Elements: Elements(els),
	}
}

// Update updates the axis values.
func (group *AxisGroup) Update() {
	tx, ty := detectAxis(group.X, group.Y, group.Elements)
	*group.X = *tx
	*group.Y = *ty
}

// Draw draws elements bound to this axis-group creating an axis automatically if necessary.
func (group *AxisGroup) Draw(plot *Plot, canvas Canvas) {
	tmpplot := &Plot{}
	*tmpplot = *plot

	if group.X != nil {
		tmpplot.X = group.X
	}
	if group.Y != nil {
		tmpplot.Y = group.Y
	}

	if !tmpplot.X.IsValid() || !tmpplot.Y.IsValid() {
		tmpplot.X, tmpplot.Y = detectAxis(tmpplot.X, tmpplot.Y, group.Elements)
	}

	for _, element := range group.Elements {
		element.Draw(tmpplot, canvas.Context(canvas.Bounds()))
	}
}
