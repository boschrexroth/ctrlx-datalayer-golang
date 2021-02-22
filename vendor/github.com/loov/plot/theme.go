package plot

import "image/color"

// Style represents a drawing style for an element.
type Style struct {
	Stroke color.Color
	Fill   color.Color
	Size   Length

	// line only
	Dash       []Length
	DashOffset []Length

	// text only
	Font     string
	Rotation float64
	Origin   Point // {-1..1, -1..1}

	// SVG
	Class string
}

// mustExists checks whether style is valid and panics if it is not.
func (style *Style) mustExist() {
	if style == nil {
		panic("style missing")
	}
}

// IsZero checks whether style has been assigned.
func (style *Style) IsZero() bool {
	if style == nil {
		return true
	}

	return style.Stroke == nil && style.Fill == nil && style.Size == 0
}

// Theme is a collection of different default styles.
type Theme struct {
	Line      Style
	Font      Style
	FontSmall Style
	Fill      Style
	Bar       Style

	Grid GridTheme
}

// GridTheme is a default style for grid.
type GridTheme struct {
	Fill  color.Color
	Major color.Color
	Minor color.Color
}

// IsZero checks whether theme has been defined.
func (theme *GridTheme) IsZero() bool {
	if theme == nil {
		return true
	}
	return theme.Fill == nil && theme.Major == nil && theme.Minor == nil
}

// NewTheme creates a theme with default values.
func NewTheme() Theme {
	return Theme{
		Line: Style{
			Stroke: color.NRGBA{0, 0, 0, 255},
			Fill:   nil,
			Size:   1.0,
		},
		Font: Style{
			Stroke: nil,
			Fill:   color.NRGBA{0, 0, 0, 255},
			Size:   12,
		},
		FontSmall: Style{
			Stroke: nil,
			Fill:   color.NRGBA{0, 0, 0, 255},
			Size:   10,
		},
		Fill: Style{
			Stroke: nil,
			Fill:   color.NRGBA{255, 255, 255, 255},
			Size:   1.0,
		},
		Bar: Style{
			Stroke: color.NRGBA{0, 0, 0, 255},
			Fill:   color.NRGBA{0, 0, 0, 100},
			Size:   1.0,
		},
		Grid: GridTheme{
			Fill:  color.NRGBA{230, 230, 230, 255},
			Major: color.NRGBA{255, 255, 255, 255},
			Minor: color.NRGBA{255, 255, 255, 100},
		},
	}
}
