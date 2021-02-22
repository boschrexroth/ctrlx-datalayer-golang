package plot

import (
	"math"
	"sort"
)

// Density implements density plot using cubic-pulse kernel.
type Density struct {
	Style
	Label string

	Kernel     Length
	Normalized bool
	Data       []float64 // sorted
}

// NewDensity creates a density plot from the given values.
func NewDensity(label string, values []float64) *Density {
	data := append(values[:0:0], values...)
	sort.Float64s(data)
	return &Density{
		Kernel:     math.NaN(),
		Label:      label,
		Normalized: true,
		Data:       data,
	}
}

// Stats calculates statistics of values.
func (line *Density) Stats() Stats {
	min, avg, max := math.NaN(), math.NaN(), math.NaN()

	n := len(line.Data)
	if n > 0 {
		min = line.Data[0]
		avg = line.Data[n/2]
		max = line.Data[n-1]
	}

	return Stats{
		Min:    Point{min, 0},
		Center: Point{avg, 0.5},
		Max:    Point{max, 1},
	}
}

// Draw draws the element to canvas.
func (line *Density) Draw(plot *Plot, canvas Canvas) {
	x, y := plot.X, plot.Y

	size := canvas.Bounds().Size()

	xmin, xmax := x.ToCanvas(x.Min, 0, size.X), x.ToCanvas(x.Max, 0, size.X)
	if xmin > xmax {
		xmin, xmax = xmax, xmin
	}

	kernel := line.Kernel
	if math.IsNaN(kernel) {
		// default to 4px wide kernel
		kernel = 4 * (x.Max - x.Min) / size.X
	}
	invkernel := 1 / kernel

	points := []Point{}
	if line.Fill != nil {
		points = append(points, Point{xmin, 0})
	}

	index := 0
	previousLow := math.Inf(-1)
	maxy := 0.0
	for screenX := 0.0; screenX < size.X; screenX += 0.5 {
		center := x.FromCanvas(screenX, 0, size.X)
		low, high := center-kernel, center+kernel

		if low < previousLow {
			index = sort.SearchFloat64s(line.Data, low)
		} else {
			for ; index < len(line.Data); index++ {
				if line.Data[index] >= low {
					break
				}
			}
		}
		previousLow = low

		sample := 0.0
		for _, value := range line.Data[index:] {
			if value > high {
				break
			}
			sample += cubicPulse(center, kernel, invkernel, value)
		}

		maxy = math.Max(maxy, sample)

		points = append(points, Point{
			X: screenX,
			Y: sample,
		})
	}

	if line.Fill != nil {
		points = append(points,
			Point{xmax, 0},
			Point{xmin, 0},
		)
	}

	scale := kernel / float64(len(line.Data))
	if line.Normalized {
		scale = 1 / maxy
	}

	for i := range points {
		points[i].Y = y.ToCanvas(points[i].Y*scale, 0, size.Y)
	}

	if !line.Style.IsZero() {
		canvas.Poly(points, &line.Style)
	} else {
		canvas.Poly(points, &plot.Theme.Line)
	}
}
