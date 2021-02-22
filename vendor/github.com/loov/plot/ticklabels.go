package plot

// TickLabels implements drawing tick labels.
type TickLabels struct {
	X *TickLabelsX
	Y *TickLabelsY
}

// NewTickLabels creates a new tick labelling element.
func NewTickLabels() *TickLabels {
	return &TickLabels{
		X: NewTickLabelsX(),
		Y: NewTickLabelsY(),
	}
}

// Draw draws tick labels to canvas using axes from plot.
func (labels *TickLabels) Draw(plot *Plot, canvas Canvas) {
	labels.X.Draw(plot, canvas)
	labels.Y.Draw(plot, canvas)
}

// TickLabelsX implements drawing tick labels.
type TickLabelsX struct {
	Enabled bool
	// Side determines position of the labels. -1 = min, 0 = center, 1 = max
	Side  float64
	Style Style
}

// NewTickLabelsX creates a new tick labelling element for X axis.
func NewTickLabelsX() *TickLabelsX {
	labels := &TickLabelsX{}
	labels.Enabled = true
	labels.Side = -1
	return labels
}

// Draw draws tick labels to canvas using axes from plot.
func (labels *TickLabelsX) Draw(plot *Plot, canvas Canvas) {
	if !labels.Enabled {
		return
	}

	x, y := plot.X, plot.Y

	sz := canvas.Bounds().Size()
	yval := lerpUnit(labels.Side, y.Min, y.Max)
	ypos := y.ToCanvas(yval, 0, sz.Y)

	style := &labels.Style
	if style.IsZero() {
		style = &plot.Theme.FontSmall
	}

	for _, tick := range x.Ticks.Ticks(x) {
		p := x.ToCanvas(tick.Value, 0, sz.X)
		if tick.Label != "" {
			canvas.Text(tick.Label, P(p, ypos), style)
		}
	}
}

// TickLabelsY implements drawing tick labels.
type TickLabelsY struct {
	Enabled bool
	// Side determines position of the labels. -1 = min, 0 = center, 1 = max
	Side  float64
	Style Style
}

// NewTickLabelsY creates a new tick labelling element for X axis.
func NewTickLabelsY() *TickLabelsY {
	labels := &TickLabelsY{}
	labels.Enabled = true
	labels.Side = -1
	return labels
}

// Draw draws tick labels to canvas using axes from plot.
func (labels *TickLabelsY) Draw(plot *Plot, canvas Canvas) {
	if !labels.Enabled {
		return
	}

	x, y := plot.X, plot.Y

	sz := canvas.Bounds().Size()
	xval := lerpUnit(labels.Side, x.Min, x.Max)
	xpos := x.ToCanvas(xval, 0, sz.X)

	style := &labels.Style
	if style.IsZero() {
		style = &plot.Theme.FontSmall
	}

	for _, tick := range y.Ticks.Ticks(y) {
		p := y.ToCanvas(tick.Value, 0, sz.Y)
		if tick.Label != "" {
			canvas.Text(tick.Label, P(xpos, p), style)
		}
	}
}
