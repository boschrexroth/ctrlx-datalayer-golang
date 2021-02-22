package plot

// Textbox implements drawing text on plots.
type Textbox struct {
	Margin Rect
	Style
	Lines []string
}

// NewTextbox creates a new text graphic with the specified lines.
func NewTextbox(line ...string) *Textbox {
	return &Textbox{
		Lines: line,
	}
}

// Add adds a line to the text box.
func (box *Textbox) Add(text string) {
	box.Lines = append(box.Lines, text)
}

// Draw draws text to the canvas.
func (box *Textbox) Draw(plot *Plot, canvas Canvas) {
	canvas = canvas.Clip(canvas.Bounds().Inset(box.Margin))
	style := box.Style
	if style.IsZero() {
		style = plot.Theme.Font
	}
	if style.Size == 0 {
		style.Size = 10
	}

	lineHeight := style.Size * 1.1
	at := Point{0, lineHeight}
	for _, line := range box.Lines {
		canvas.Text(line, at, &style)
		at.Y += lineHeight
	}
}
