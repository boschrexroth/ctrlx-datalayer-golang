package plot

import (
	"math"
	"sort"
)

// Percentiles implements drawing percentile line.
type Percentiles struct {
	Style
	Label string
	Data  []Point
}

// NewPercentiles creates percentiles from values.
func NewPercentiles(label string, values []float64) *Percentiles {
	values = append(values[:0:0], values...)
	sort.Float64s(values)

	points := make([]Point, 0, len(values))

	multiplier := 1 / float64(len(values)+1)
	for i, v := range values {
		var p Point
		p.X = float64(i+1) * multiplier
		p.Y = v
		points = append(points, p)
	}

	return &Percentiles{
		Label: label,
		Data:  points,
	}
}

// Stats calculates element statistics.
func (line *Percentiles) Stats() Stats {
	return PointsStats(line.Data)
}

// Draw draws the element to canvas.
func (line *Percentiles) Draw(plot *Plot, canvas Canvas) {
	canvas = canvas.Clip(canvas.Bounds())

	points := make([]Point, 0, len(line.Data))
	lastScreenX := math.Inf(-1)
	projectcb(line.Data, plot.X, plot.Y, canvas.Bounds(), func(p Point) {
		if math.Abs(lastScreenX-p.X) > 0.5 {
			points = append(points, p)
			lastScreenX = p.X
		}
	})

	if !line.Style.IsZero() {
		canvas.Poly(points, &line.Style)
	} else {
		canvas.Poly(points, &plot.Theme.Line)
	}
}

// PercentilesTransform implements axis transform for percentiles X axis.
type PercentileTransform struct {
	levels  int
	base    float64
	mulbase float64
}

// NewPercentileTransform creates a transform for X axis with the specified max levels.
func NewPercentileTransform(levels int) *PercentileTransform {
	base := math.Pow(0.1, float64(levels))
	return &PercentileTransform{
		levels:  levels,
		base:    base,
		mulbase: 1 / math.Log(base),
	}
}

// transform transforms axis range to normalized value.
func (tx *PercentileTransform) transform(v float64) float64 {
	return math.Log(1-v) * tx.mulbase
}

// inverse transform normalized value to axis range.
func (tx *PercentileTransform) inverse(v float64) float64 {
	return 1 - math.Pow(tx.base, v)
}

// ToCanvas converts value v to canvas space.
func (tx *PercentileTransform) ToCanvas(axis *Axis, v float64, screenMin, screenMax Length) Length {
	v = tx.transform(v)
	low, high := axis.lowhigh()
	n := (v - low) / (high - low)
	r := screenMin + n*(screenMax-screenMin)

	// when v == 0, we do not want to return an infinity
	if r > screenMax {
		return screenMax
	}
	return r
}

// FromCanvas converts screen value s to value.
func (tx *PercentileTransform) FromCanvas(axis *Axis, s Length, screenMin, screenMax Length) float64 {
	low, high := axis.lowhigh()
	n := (s - screenMin) / (screenMax - screenMin)
	v := low + n*(high-low)
	return tx.inverse(v)
}

// NewPercentilesAxis creates X axis for plotting percentiles.
// It uses 5 digits of precision max.
func NewPercentilesAxis() *Axis {
	axis := NewAxis()
	axis.Transform = NewPercentileTransform(5)
	axis.Ticks = ManualTicks{
		{Value: 0, Label: "0"},
		{Value: 0.25, Label: "25"},
		{Value: 0.5, Label: "50"},
		{Value: 0.75, Label: "75"},
		{Value: 0.9, Label: "90"},
		{Value: 0.99, Label: "99"},
		{Value: 0.999, Label: "99.9"},
		{Value: 0.9999, Label: "99.99"},
		{Value: 0.99999, Label: "99.999"}}
	return axis
}
