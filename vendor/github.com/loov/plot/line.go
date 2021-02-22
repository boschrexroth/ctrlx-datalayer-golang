package plot

import "math"

// Line implements a simple line plot.
type Line struct {
	Style
	Label string

	Data []Point
}

// NewLine creates a new line element from the given points.
func NewLine(label string, points []Point) *Line {
	return &Line{
		Label: label,
		Data:  points,
	}
}

// Stats calculates element statistics.
func (line *Line) Stats() Stats {
	return PointsStats(line.Data)
}

// Draw draws the element to canvas.
func (line *Line) Draw(plot *Plot, canvas Canvas) {
	canvas = canvas.Clip(canvas.Bounds())
	points := project(line.Data, plot.X, plot.Y, canvas.Bounds())

	if !line.Style.IsZero() {
		canvas.Poly(points, &line.Style)
	} else {
		canvas.Poly(points, &plot.Theme.Line)
	}
}

// OptimizedLine implements a simple line plot.
type OptimizedLine struct {
	Style
	Label string

	Data []Point

	ThresholdPx float64
}

// NewOptimizedLine creates a new line element that tries to optimize drawing.
func NewOptimizedLine(label string, points []Point, thresholdPx float64) *OptimizedLine {
	return &OptimizedLine{
		Label:       label,
		ThresholdPx: thresholdPx,
		Data:        points,
	}
}

// Stats calculates element statistics.
func (line *OptimizedLine) Stats() Stats {
	return PointsStats(line.Data)
}

// Draw draws the element to canvas.
func (line *OptimizedLine) Draw(plot *Plot, canvas Canvas) {
	canvas = canvas.Clip(canvas.Bounds())
	points := project(line.Data, plot.X, plot.Y, canvas.Bounds())

	const optimizeCount = 100
	if len(points) < optimizeCount {
		if !line.Style.IsZero() {
			canvas.Poly(points, &line.Style)
		} else {
			canvas.Poly(points, &plot.Theme.Line)
		}
		return
	}

	// always include the first point
	optimized := points[:1]

	prev := points[0]
	mid := points[1]
	for _, next := range points[2:] {
		// does the mid change significantly from the previous?
		if math.Abs(prev.Y-mid.Y) < line.ThresholdPx && math.Abs(prev.X-mid.X) < line.ThresholdPx {
			mid = next
			continue
		}

		// is the mid on the line from prev to next?
		p := invlerp(mid.X, prev.X, next.X)
		knownY := lerp(p, prev.Y, next.Y)
		if math.Abs(knownY-mid.X) < line.ThresholdPx {
			mid = next
			continue
		}

		// otherwise let's output
		optimized = append(optimized, mid)
		prev, mid = mid, next
	}
	// add the final point
	optimized = append(optimized, points[len(points)-1])

	if !line.Style.IsZero() {
		canvas.Poly(optimized, &line.Style)
	} else {
		canvas.Poly(optimized, &plot.Theme.Line)
	}
}
