package sio

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

// Text Drawing Status
var (
	DefaultEmHeight   = 12.0
	DefaultEmWidth    = 6.0
	DefaultLineHeight = 1.3
)

type TextBox struct {
	Anchor            int
	Rect              *Rect
	EmHeight, EmWidth float64
	LineHeight        float64

	text string
}

func (r *Rect) NewTextBox(text string, anchor int) *TextBox {
	return &TextBox{
		Anchor:     anchor,
		Rect:       r,
		EmHeight:   DefaultEmHeight,
		EmWidth:    DefaultEmWidth,
		LineHeight: DefaultLineHeight,
		text:       text,
	}
}

type TextRow struct {
	Text string
	X, Y int
}

func TextRows(text string, re *Rect) []TextRow {
	rows := []TextRow{}

	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return rows
	}

	// make vertical ruler
	h := float64(len(lines)-1) * DefaultEmHeight * DefaultLineHeight
	h += DefaultEmHeight
	x, y := re.Pos(re.anchor)
	vert := NewRect(re.anchor, x, y, 0, h)
	x, y = vert.Pos(7) // top of the vert

	h = 0
	for _, line := range lines {
		w := runewidth.StringWidth(line)
		horz := NewRect(re.anchor, x, y+h, float64(w)*DefaultEmWidth, 0)
		u, v := horz.Pos(7)

		rows = append(rows, TextRow{
			Text: line,
			X:    int(u + 0.5),
			Y:    int(v + 0.5),
		})
	}

	return rows
}
