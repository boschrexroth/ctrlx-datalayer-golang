package plot

import "time"

// DurationTo converts durations to the specified scale.
func DurationTo(durations []time.Duration, scale time.Duration) []float64 {
	values := make([]float64, len(durations))
	for i, dur := range durations {
		values[i] = float64(dur) / float64(scale)
	}
	return values
}

// DurationToNanoseconds converts an slice of durations to nanoseconds.
func DurationToNanoseconds(durations []time.Duration) []float64 {
	return DurationTo(durations, time.Nanosecond)
}

// DurationToSeconds converts an slice of durations to seconds.
func DurationToSeconds(durations []time.Duration) []float64 {
	return DurationTo(durations, time.Second)
}
