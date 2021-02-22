package plot

// Label describes a plot element that draws a text to relative position.
type Label struct {
	Placement Point
	Style
	Text string
}

// NewXLabel creates a text label that is placed according to relative placement.
func NewXLabel(text string) *Label {
	return &Label{
		Placement: Point{0, 1},
		Style: Style{
			Origin: Point{0, -1},
		},
		Text: text,
	}
}

// Draw draws the element to canvas.
func (label *Label) Draw(plot *Plot, canvas Canvas) {
	style := &label.Style
	if style == nil {
		t := plot.Theme.Font
		t.Origin = label.Style.Origin
		style = &t
	}

	bounds := canvas.Bounds()
	at := bounds.UnitLocation(label.Placement)

	canvas.Text(label.Text, at, style)
}
