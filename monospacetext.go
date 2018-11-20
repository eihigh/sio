package sio

import "github.com/mattn/go-runewidth"

// Text Drawing Statuses
var (
	DefaultEmHeight = 12.0
	DefaultEmWidth  = 6.0

	DefaultLineHeight = 1.5
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

// NewTextBox creates an instance with default profile
func NewTextBox(anchor int, text string) *TextBox {
	return &TextBox{
		Anchor:     anchor,
		Rect:       NewRect(7, 0, 0, 0, 0),
		EmHeight:   DefaultEmHeight,
		EmWidth:    DefaultEmWidth,
		LineHeight: DefaultLineHeight,
		text:       []rune(text),
	}
}

// SetText sets the text
func (t *TextBox) SetText(text string) {
	t.text = []rune(text)
	t.pos = 0
}

// SeekRune adds the last text position.
func (t *TextBox) SeekRune() {
	t.pos++
}

// SeekLine adds the line.
func (t *TextBox) SeekLine() {
	for {
		t.pos++
		if t.text[t.pos] == '\n' {
			return
		}
	}
}

// SeekAll seeks to the end
func (t *TextBox) SeekAll() {
	t.pos = len(t.text)
}

// TextLineInfo contains info to draw a line of text
type TextLineInfo struct {
	Text string
	X, Y int
}

// Lines returns []TextLineInfo
func (t *TextBox) Lines() []TextLineInfo {

	lines := make([]TextLineInfo, 0)

	end := len(t.text)
	if t.IgnoreHiddenText {
		end = t.pos
	}

	// get end of each lines
	lns := make([]int, 0)
	for n, r := range t.text[:end] {
		if r == '\n' {
			lns = append(lns, n)
		}
	}

	// make ruler (zero width)
	h := float64(len(lns)) * t.EmHeight
	h += float64(len(lns[1:])) * t.EmHeight * (t.LineHeight - 1.0)
	x, y := t.Rect.Pos(t.Anchor)
	horz := NewRect(t.Anchor, x, y, 0, h)

	// split lines and get runewidths
	n := 0
	h = 0
	for _, m := range lns {
		w := 0
		for _, r := range t.text[n:m] {
			w += runewidth.RuneWidth(r)
		}
		x, y = horz.Pos(7) // from (left-)top
		vert := NewRect(t.Anchor, x, y+h, float64(w)*t.EmWidth, 0)
		x, y = vert.Pos(7)
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
