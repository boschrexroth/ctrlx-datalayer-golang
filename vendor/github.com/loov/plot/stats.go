package plot

import "math"

// Stats describes a data points.
type Stats struct {
	Min    Point
	Center Point
	Max    Point
}

// nanStats is used for missing stats.
var nanStats = Stats{
	Min:    nanPoint,
	Center: nanPoint,
	Max:    nanPoint,
}

// tryGetStats tries to calculate stats.
func tryGetStats(element Element) (Stats, bool) {
	if dataset, ok := element.(Dataset); ok {
		return dataset.Stats(), true
	}
	return nanStats, false
}

// maximalStats finds bounding stats from elements.
func maximalStats(els []Element) Stats {
	finalstats := nanStats
	for _, element := range els {
		if element == nil {
			continue
		}

		if elstats, ok := tryGetStats(element); ok {
			if math.IsNaN(finalstats.Min.X) {
				finalstats.Min.X = elstats.Min.X
			} else if !math.IsNaN(elstats.Min.X) {
				finalstats.Min.X = math.Min(finalstats.Min.X, elstats.Min.X)
			}
			if math.IsNaN(finalstats.Max.X) {
				finalstats.Max.X = elstats.Max.X
			} else if !math.IsNaN(elstats.Max.X) {
				finalstats.Max.X = math.Max(finalstats.Max.X, elstats.Max.X)
			}

			if math.IsNaN(finalstats.Min.Y) {
				finalstats.Min.Y = elstats.Min.Y
			} else if !math.IsNaN(elstats.Min.Y) {
				finalstats.Min.Y = math.Min(finalstats.Min.Y, elstats.Min.Y)
			}
			if math.IsNaN(finalstats.Max.Y) {
				finalstats.Max.Y = elstats.Max.Y
			} else if !math.IsNaN(elstats.Max.Y) {
				finalstats.Max.Y = math.Max(finalstats.Max.Y, elstats.Max.Y)
			}
		}
	}
	return finalstats
}

// PointsStats calculates statistics of points.
func PointsStats(points []Point) Stats {
	min, avg, max := nanPoint, Point{}, nanPoint

	if len(points) > 0 {
		min = points[0]
		max = points[0]
	}

	for _, p := range points {
		min = min.Min(p)
		avg = avg.Add(p)
		max = max.Max(p)
	}

	return Stats{
		Min:    min,
		Center: avg.Scale(1 / float64(len(points))),
		Max:    max,
	}
}
