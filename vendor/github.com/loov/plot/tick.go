package plot

import (
	"fmt"
	"math"
)

// Ticks represents an approach to calculating tick position and label.
type Ticks interface {
	Ticks(axis *Axis) []Tick
}

// Tick represents a division on plot.
type Tick struct {
	Minor bool
	Label string
	Value float64
}

// AutomaticTicks tries to automatically figure out which ticks to use.
type AutomaticTicks struct{}

// logarithmicTicks calculates ticks for logarithmic axis.
func (at AutomaticTicks) logarithmicTicks(axis *Axis, transform *Log1pTransform) []Tick {
	//TODO: fix, we don't properly assign labels for logarithmic axis

	ticks := make([]Tick, 0)

	low, high := axis.Min, axis.Max
	if low > high {
		low, high = high, low
	}

	previous := math.NaN()

	inRange := func(value float64) bool {
		return low < value && value < high
	}

	if inRange(0) {
		ticks = append(ticks, Tick{Value: 0, Label: "0"})
		previous = 0
	}

	for power := 0; power < 10; power++ {
		value := math.Pow(transform.base, float64(power))
		if inRange(value) {
			ticks = append(ticks, Tick{
				Value: value,
				Label: fmt.Sprintf("%.0f", value),
			})
		}
		if inRange(-value) {
			ticks = append(ticks, Tick{
				Value: -value,
				Label: fmt.Sprintf("%.0f", -value),
			})
		}

		if !math.IsNaN(previous) && axis.MinorTicks > 0 {
			minorSpacing := (value - previous) / float64(axis.MinorTicks)
			minor := previous
			for i := 0; i < axis.MinorTicks; i++ {
				if inRange(minor) {
					ticks = append(ticks, Tick{Minor: true, Value: minor})
				}
				if inRange(-minor) {
					ticks = append(ticks, Tick{Minor: true, Value: -minor})
				}
				minor += minorSpacing
			}
		}
		previous = value
	}

	return ticks
}

// linearTicks calculates ticks on linear axis.
func (AutomaticTicks) linearTicks(axis *Axis) []Tick {
	majorSpacing := (axis.Max - axis.Min) / float64(axis.MajorTicks)
	minorSpacing := majorSpacing / float64(axis.MinorTicks)

	frac := -int(math.Floor(math.Log10(majorSpacing)))
	if frac < 0 {
		frac = 0
	}

	ticks := make([]Tick, 0, axis.MajorTicks*axis.MinorTicks)

	major := axis.Min
	hasZero := false
	for i := 0; i < axis.MajorTicks; i++ {
		if major == 0 {
			hasZero = true
		}
		ticks = append(ticks, Tick{
			Value: major,
			Label: fmt.Sprintf("%.[2]*[1]f", major, frac),
		})

		minor := major
		for k := 0; k < axis.MinorTicks; k++ {
			if minor == 0 {
				minor += minorSpacing
				continue
			}
			ticks = append(ticks, Tick{
				Minor: true,
				Value: minor,
			})
			minor += minorSpacing
		}

		major += majorSpacing
	}

	if !hasZero && (axis.Min <= 0) == (0 <= axis.Max) {
		ticks = append(ticks, Tick{
			Value: 0,
			Label: "0",
		})
	}

	return ticks
}

// Ticks automatically calculates appropriate ticks for an axis.
func (ticks AutomaticTicks) Ticks(axis *Axis) []Tick {
	// if transform, ok := axis.Transform.(*Log1pTransform); ok {
	// 	return ticks.logarithmicTicks(axis, transform)
	// }
	return ticks.linearTicks(axis)
}

// ManualTicks allows to manually place and label ticks.
type ManualTicks []Tick

// Ticks calculates ticks for specified axis.
func (ticks ManualTicks) Ticks(axis *Axis) []Tick { return []Tick(ticks) }
