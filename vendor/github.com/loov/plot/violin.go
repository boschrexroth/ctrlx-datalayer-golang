package plot

import (
	"math"
	"sort"
)

// Violin implements violin plot using cubic-pulse for kernel function.
type Violin struct {
	Style
	Label string

	Side       float64
	Kernel     Length
	Normalized bool
	Data       []float64 // sorted
}

// NewViolin creates a new violin element using the specified values.
func NewViolin(label string, values []float64) *Violin {
	data := append(values[:0:0], values...)
	sort.Float64s(data)
	return &Violin{
		Kernel:     math.NaN(),
		Side:       1,
		Label:      label,
		Normalized: true,
		Data:       data,
	}
}

// Stats calculates element statistics.
func (line *Violin) Stats() Stats {
	min, avg, max := math.NaN(), math.NaN(), math.NaN()

	n := len(line.Data)
	if n > 0 {
		min = line.Data[0]
		avg = line.Data[n/2]
		max = line.Data[n-1]
	}

	return Stats{
		Min:    Point{-1, min},
		Center: Point{0, avg}, // todo, figure out how to get the 50% of density plot
		Max:    Point{1, max},
	}
}

// Draw draws the element to canvas.
func (line *Violin) Draw(plot *Plot, canvas Canvas) {
	x, y := plot.X, plot.Y

	size := canvas.Bounds().Size()

	ymin, ymax := y.ToCanvas(y.Min, 0, size.Y), y.ToCanvas(y.Max, 0, size.Y)
	if ymin > ymax {
		ymin, ymax = ymax, ymin
	}

	kernel := line.Kernel
	if math.IsNaN(kernel) {
		// default to 4px wide kernel
		kernel = 4 * (y.Max - y.Min) / size.Y
	}
	invkernel := 1 / kernel

	points := []Point{}
	if line.Fill != nil || line.Side == 0 {
		points = append(points, Point{0, ymin})
	}

	index := 0
	previousLow := math.Inf(-1)
	maxx := 0.0
	for screenY := 0.0; screenY < size.Y; screenY += 0.5 {
		center := y.FromCanvas(screenY, 0, size.Y)
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
		maxx = math.Max(maxx, sample)

		points = append(points, Point{
			X: sample,
			Y: screenY,
		})
	}

	if line.Fill != nil || line.Side == 0 {
		points = append(points, Point{0, ymax})
	}

	scale := kernel / float64(len(line.Data))
	if line.Normalized {
		scale = 1 / maxx
	}

	if line.Side == 0 {
		otherSide := make([]Point, len(points))
		for i := range points {
			k := len(points) - i - 1
			otherSide[k] = points[i]
			points[i].X = x.ToCanvas(points[i].X*scale, 0, size.X)
			otherSide[k].X = x.ToCanvas(-otherSide[k].X*scale, 0, size.X)
		}
		points = append(points, otherSide...)
	} else {
		scale *= line.Side
		for i := range points {
			points[i].X = x.ToCanvas(points[i].X*scale, 0, size.X)
		}
	}

	if !line.Style.IsZero() {
		canvas.Poly(points, &line.Style)
	} else {
		canvas.Poly(points, &plot.Theme.Line)
	}
}
