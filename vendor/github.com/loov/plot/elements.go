package plot

// Elements defines a list of elements.
type Elements []Element

// Add appends an element to the list.
func (els *Elements) Add(el Element) { *els = append(*els, el) }

// AddGroup appends elements as a single group to the list.
func (els *Elements) AddGroup(adds ...Element) { els.Add(Elements(adds)) }

// Stats calculates the stats from all elements.
func (els Elements) Stats() Stats { return maximalStats(els) }

// Draw draws the elements drawn over each other.
func (els Elements) Draw(plot *Plot, canvas Canvas) {
	for _, el := range els {
		if el == nil {
			continue
		}
		el.Draw(plot, canvas)
	}
}

// Margin is a collection which is drawn with a margin.
type Margin struct {
	Amount Rect
	Elements
}

// NewMargin creates a elements groups.
func NewMargin(amount Rect, els ...Element) *Margin {
	return &Margin{Amount: amount, Elements: Elements(els)}
}

// Draw draws the elements drawn over each other.
func (margin *Margin) Draw(plot *Plot, canvas Canvas) {
	bounds := canvas.Bounds().Inset(margin.Amount)
	margin.Elements.Draw(plot, canvas.Context(bounds))
}

// VStack implements vertically stacked elements.
type VStack struct {
	Margin Rect
	Elements
}

// NewVStack creates a collection of elements that are vertically stacked.
func NewVStack(els ...Element) *VStack {
	return &VStack{Elements: Elements(els)}
}

// Draw draws elements vertically stacked equally dividing space.
func (stack *VStack) Draw(plot *Plot, canvas Canvas) {
	if len(stack.Elements) == 0 {
		return
	}
	bounds := canvas.Bounds()
	for i, el := range stack.Elements {
		block := bounds.Row(i, len(stack.Elements))
		el.Draw(plot, canvas.Context(block.Inset(stack.Margin)))
	}
}

// HStack implements horizontally stacked elements.
type HStack struct {
	Margin Rect
	Elements
}

// NewHStack creates a collection of elements that are horizontally stacked.
func NewHStack(els ...Element) *HStack {
	return &HStack{Elements: Elements(els)}
}

// Draw draws elements horizontally stacked equally dividing space.
func (stack *HStack) Draw(plot *Plot, canvas Canvas) {
	if len(stack.Elements) == 0 {
		return
	}
	bounds := canvas.Bounds()
	for i, el := range stack.Elements {
		block := bounds.Column(i, len(stack.Elements))
		el.Draw(plot, canvas.Context(block.Inset(stack.Margin)))
	}
}

// HFlex implements horizontally stacked elements with non-equal sizes.
type HFlex struct {
	Margin Rect

	fixedSize []float64
	elements  Elements
}

// NewHFlex creates a horizontally flexing elements.
func NewHFlex() *HFlex { return &HFlex{} }

// Stats calculates the stats from all elements.
func (stack *HFlex) Stats() Stats { return stack.elements.Stats() }

// Add adds an element with fixed size.
func (stack *HFlex) Add(fixedSize float64, el Element) {
	stack.elements.Add(el)
	stack.fixedSize = append(stack.fixedSize, fixedSize)
}

// AddGroup adds a group of elements with fixed size.
func (stack *HFlex) AddGroup(fixedSize float64, adds ...Element) {
	stack.Add(fixedSize, Elements(adds))
}

// Draw draws elements.
func (stack *HFlex) Draw(plot *Plot, canvas Canvas) {
	if len(stack.elements) == 0 {
		return
	}

	fixedSize := 0.0
	flexCount := 0.0
	for i, size := range stack.fixedSize {
		fixedSize += size
		if stack.elements[i] == nil {
			continue
		}
		if size == 0 {
			flexCount++
		}
	}

	bounds := canvas.Bounds()
	size := bounds.Size()

	flexWidth := (bounds.Size().X - fixedSize) / flexCount
	min := bounds.Min
	for i, el := range stack.elements {
		elsize := stack.fixedSize[i]
		if el == nil {
			min.X += elsize
			continue
		}
		if elsize == 0 {
			elsize = flexWidth
		}

		block := Rect{
			min,
			min.Add(Point{elsize, size.Y}),
		}
		min.X = block.Max.X

		el.Draw(plot, canvas.Context(block.Inset(stack.Margin)))
	}
}
