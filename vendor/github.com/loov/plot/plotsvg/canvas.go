package plotsvg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sort"

	"github.com/loov/plot"
)

var _ plot.Canvas = (*Canvas)(nil)

// Canvas describes the top-level svg drawing context.
type Canvas struct {
	Style string
	context
}

// New creates a new SVG canvas.
func New(width, height plot.Length) *Canvas {
	svg := &Canvas{}
	svg.Style = `text { text-shadow: -1px -1px 0 rgba(255,255,255,0.5),	1px -1px 0 rgba(255,255,255,0.5), 1px  1px 0 rgba(255,255,255,0.5), -1px  1px 0 rgba(255,255,255,0.5); }`
	svg.bounds.Max.X = width
	svg.bounds.Max.Y = height
	return svg
}

// context describes a svg drawing context.
type context struct {
	index int
	clip  bool
	// bounds relative to parent
	bounds   plot.Rect
	elements []element
	layers   []*context
}

// element describes an svg element.
type element struct {
	// style
	style plot.Style
	// line
	points []plot.Point
	// text
	text   string
	origin plot.Point
	// context
	context *context
}

// Bytes returns svg context as a byte array.
func (svg *Canvas) Bytes() []byte {
	var buffer bytes.Buffer
	svg.WriteTo(&buffer)
	return buffer.Bytes()
}

// Bounds returns the bounds in the global size.
func (svg *context) Bounds() plot.Rect {
	return svg.bounds.Zero()
}

// Size returns the size of the context.
func (svg *context) Size() plot.Point {
	return svg.bounds.Size()
}

// context creates a clipped subcontext.
func (svg *context) context(r plot.Rect, clip bool) plot.Canvas {
	element := element{}
	element.context = &context{}
	element.context.clip = clip
	element.context.bounds = r
	svg.elements = append(svg.elements, element)
	return element.context
}

// context creates a subcontext bounded to r.
func (svg *context) Context(r plot.Rect) plot.Canvas {
	return svg.context(r, false)
}

// Clip clips to rect.
func (svg *context) Clip(r plot.Rect) plot.Canvas {
	return svg.context(r, true)
}

// Layer returns an layer above or below current state.
func (svg *context) Layer(index int) plot.Canvas {
	if index == 0 {
		return svg
	}

	i := sort.Search(len(svg.layers), func(i int) bool {
		return svg.layers[i].index > index
	})
	if i < len(svg.layers) && svg.layers[i].index == index {
		return svg.layers[i]
	} else {
		layer := &context{}
		layer.index = index
		layer.bounds = svg.bounds.Zero()

		svg.layers = append(svg.layers, layer)
		copy(svg.layers[i+1:], svg.layers[i:])
		svg.layers[i] = layer
		return layer
	}
}

// Text draws text.
func (svg *context) Text(text string, at plot.Point, style *plot.Style) {
	mustExist(style)
	svg.elements = append(svg.elements, element{
		text:   text,
		origin: at,
		style:  *style,
	})
}

// Poly draws a polygon.
func (svg *context) Poly(points []plot.Point, style *plot.Style) {
	mustExist(style)
	svg.elements = append(svg.elements, element{
		points: points,
		style:  *style,
	})
}

// Rect draws a rectangle.
func (svg *context) Rect(r plot.Rect, style *plot.Style) {
	svg.Poly(r.Points(), style)
}

// WriteTo writes svg content to dst.
func (svg *Canvas) WriteTo(dst io.Writer) (n int64, err error) {
	w := &writer{}
	w.Writer = dst

	id := 0

	// svg wrapper
	w.Print(`<?xml version="1.0" standalone="no"?>`)
	w.Print(`<!DOCTYPE svg PUBLIC "-//W3C//DTD Canvas 1.0//EN" "http://www.w3.org/TR/2001/REC-Canvas-20010904/DTD/svg10.dtd">`)
	size := svg.bounds.Size()
	w.Print(`<svg xmlns='http://www.w3.org/2000/svg' xmlns:loov='http://www.loov.io' width='%vpx' height='%vpx'>`, size.X, size.Y)
	defer w.Print(`</svg>`)

	if svg.Style != "" {
		// TODO: escape CDATA
		w.Print(`<style>/* <![CDATA[ */ %v /* ]]> */ </style>`, svg.Style)
	}

	w.Print(`<g transform='translate(0.5, 0.5)'>`)
	defer w.Print(`</g>`)

	var writeLayer func(svg *context)
	var writeElement func(svg *context, el *element)

	writeLayer = func(svg *context) {
		if svg.clip {
			id++
			size := svg.bounds.Size()
			w.Print(`<clipPath id='clip%v'><rect x='0' y='0' width='%v' height='%v' /></clipPath>`, id, size.X, size.Y)
		}

		w.Printf(`<g`)
		w.Printf(` loov:index='%v'`, svg.index)
		if svg.bounds.Min.X != 0 || svg.bounds.Min.Y != 0 {
			w.Printf(` transform='translate(%.2f %.2f)'`, svg.bounds.Min.X, svg.bounds.Min.Y)
		}
		if svg.clip {
			w.Printf(` clip-path='url(#clip%v)'`, id)
		}

		w.Print(">")
		defer w.Print(`</g>`)

		after := 0
		for i, layer := range svg.layers {
			if layer.index >= 0 {
				after = i
				break
			}
			writeLayer(layer)
		}

		if len(svg.elements) > 0 {
			if len(svg.layers) > 0 {
				w.Print("<g>")
			}
			for i := range svg.elements {
				writeElement(svg, &svg.elements[i])
			}
			if len(svg.layers) > 0 {
				w.Print("</g>")
			}
		}

		for _, layer := range svg.layers[after:] {
			writeLayer(layer)
		}
	}

	writeElement = func(svg *context, el *element) {
		if len(el.points) > 0 {
			w.Printf(`<polyline `)
			w.writePolyStyle(&el.style)
			w.Printf(` points='`)
			for _, p := range el.points {
				w.Printf(`%.2f,%.2f `, p.X, p.Y)
			}
			w.Print(`' />`)
		}
		if el.text != "" {
			w.Printf(`<text x='%.2f' y='%.2f' `, el.origin.X, el.origin.Y)
			w.writeTextStyle(&el.style)
			w.Printf(`>`)
			xml.EscapeText(w, []byte(el.text))
			w.Print(`</text>`)
		}
		if el.context != nil {
			writeLayer(el.context)
		}
	}

	writeLayer(&svg.context)

	return w.total, w.err
}

// writeTextStyle writes text styling.
func (w *writer) writeTextStyle(style *plot.Style) {
	// TODO: merge with writePolyStyle
	if style.Class != "" {
		w.Printf(` class='`)
		xml.EscapeText(w, []byte(style.Class))
		w.Printf(`'`)
	}

	if style.Font == "" && style.Size == 0 && style.Stroke == nil && style.Fill == nil {
		return
	}

	w.Printf(` style='`)
	defer w.Printf(`' `)

	if style.Font != "" {
		w.Printf(`font-family: "`)
		// TODO, verify escaping
		xml.EscapeText(w, []byte(style.Font))
		w.Printf(`";`)
	}
	if style.Size != 0 {
		w.Printf(`font-size: %vpx;`, style.Size)
	}
	if style.Stroke != nil {
		w.Printf(`color: %v;`, convertColorToHex(style.Stroke))
	}
	if style.Fill != nil {
		w.Printf(`fill: %v;`, convertColorToHex(style.Fill))
	}
}

// writePolyStyle writes a polygon class and styling.
func (w *writer) writePolyStyle(style *plot.Style) {
	if style.Class != "" {
		w.Printf(` class='`)
		xml.EscapeText(w, []byte(style.Class))
		w.Printf(`'`)
	}

	if style.Size == 0 && style.Stroke == nil && style.Fill == nil && len(style.Dash) == 0 {
		return
	}

	w.Printf(` style='`)
	defer w.Printf(`'`)

	if style.Stroke != nil {
		w.Printf(`stroke: %v;`, convertColorToHex(style.Stroke))
	} else {
		w.Printf(`stroke: transparent;`)
	}

	if style.Fill != nil {
		w.Printf(`fill: %v;`, convertColorToHex(style.Fill))
	} else {
		w.Printf(`fill: transparent;`)
	}

	if len(style.Dash) > 0 {
		w.Printf(`stroke-dasharray:`)
		for _, v := range style.Dash {
			w.Printf(` %v`, v)
		}
		w.Printf(`;`)
	}

	if style.Size != 0 {
		w.Printf(`stroke-width: %vpx;`, style.Size)
	}

	// TODO: dash
}

// writer implements error capturing/hiding writer.
type writer struct {
	io.Writer
	total int64
	err   error
}

// Errored returns whether there has been an error during writing.
func (w *writer) Errored() bool { return w.err != nil }

// Write is a convenience func for writing svg content.
func (w *writer) Write(data []byte) (int, error) {
	if w.Errored() {
		return 0, nil
	}
	n, err := w.Writer.Write(data)
	w.total += int64(n)
	if err != nil {
		w.err = err
	}
	return n, nil
}

// Print is a convenience function for writing svg content.
func (w *writer) Print(format string, args ...interface{}) { fmt.Fprintf(w, format+"\n", args...) }

// Printf is a convenience function for writing svg content.
func (w *writer) Printf(format string, args ...interface{}) { fmt.Fprintf(w, format, args...) }


// mustExists checks whether style is valid and panics if it is not.
func mustExist(style *plot.Style)  {
	if style == nil {
		panic("style missing")
	}
}
