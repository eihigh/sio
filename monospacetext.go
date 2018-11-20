package sio

import "github.com/mattn/go-runewidth"

// Text Drawing Status
var (
	DefaultEmHeight = 12.0
	DefaultEmWidth  = 6.0

	DefaultLineHeight = 1.3
)

// TextBox provides some functions for monospace font drawing
type TextBox struct {
	IgnoreHiddenText  bool
	Anchor            int
	Rect              *Rect
	EmHeight, EmWidth float64
	LineHeight        float64

	text []rune
	pos  int
}

// NewTextBox creates a TextBox with default profile.
func NewTextBox(anchor int, text string) *TextBox {
	tb := NewTypeWriter(anchor, text)
	tb.Reveal()
	return tb
}

// NewTypeWriter creates a TextBox with default profile.
func NewTypeWriter(anchor int, text string) *TextBox {
	return &TextBox{
		Anchor:     anchor,
		Rect:       NewRect(7, 0, 0, 0, 0),
		EmHeight:   DefaultEmHeight,
		EmWidth:    DefaultEmWidth,
		LineHeight: DefaultLineHeight,
		text:       []rune(text),
	}
}

// SetText sets the text.
func (t *TextBox) SetText(text string) {
	t.text = []rune(text)
	t.pos = 0
}

// NextRune makes the next rune visible.
func (t *TextBox) NextRune() {
	t.pos++
}

// NextLine makes the next line visible.
func (t *TextBox) NextLine() {
	for {
		t.pos++
		if t.text[t.pos] == '\n' {
			return
		}
	}
}

// Reveal makes all text visible.
func (t *TextBox) Reveal() {
	t.pos = len(t.text)
}

// TextLineInfo contains info to draw a line of text.
type TextLineInfo struct {
	Text string
	X, Y int
}

// Lines returns []TextLineInfo
func (t *TextBox) Lines() []TextLineInfo {

	lines := make([]TextLineInfo, 0)

	included := len(t.text)
	if t.IgnoreHiddenText {
		included = t.pos
	}

	// get line fields
	lns := make([]int, 0)
	for i, r := range t.text[:included] {
		if r == '\n' {
			lns = append(lns, i)
		}
	}

	// make ruler (zero width)
	h := float64(len(lns)) * t.EmHeight * t.LineHeight // lines with line break
	h += t.EmHeight                                    // first line
	x, y := t.Rect.Pos(t.Anchor)
	vert := NewRect(t.Anchor, x, y, 0, h)

	// make line infos
	lns = append(lns, len(t.text))
	n := 0
	h = 0
	for _, m := range lns {
		w := 0
		for _, r := range t.text[n:m] {
			w += runewidth.RuneWidth(r)
		}
		x, y = vert.Pos(7) // (left-)top position
		horz := NewRect(t.Anchor, x, y+h, float64(w)*t.EmWidth, 0)
		x, y = horz.Pos(7)
		line := TextLineInfo{
			Text: string(t.text[n:m]),
			X:    int(x + 0.5),
			Y:    int(y + 0.5),
		}
		lines = append(lines, line)

		h += t.EmHeight * t.LineHeight
		n = m + 1
	}

	return lines
}
