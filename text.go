package sio

import (
	"strings"

	"github.com/mattn/go-runewidth"
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

type TextLineInfo struct {
	Text string
	X, Y int
}

func (t *TextBox) Lines() []TextLineInfo {

	infos := []TextLineInfo{}

	lines := strings.Split(t.text, "\n")
	if len(lines) == 0 {
		return infos
	}

	// make vertical ruler
	h := float64(len(lines)-1) * t.EmHeight * t.LineHeight
	h += t.EmHeight
	x, y := t.Rect.Pos(t.Anchor)
	vert := NewRect(t.Anchor, x, y, 0, h)
	x, y = vert.Pos(7) // top of the vert

	// make line infos
	h = 0
	for _, line := range lines {
		w := runewidth.StringWidth(line)
		horz := NewRect(t.Anchor, x, y+h, float64(w)*t.EmWidth, 0)
		u, v := horz.Pos(7) // left of the horz

		infos = append(infos, TextLineInfo{
			Text: line,
			X:    int(u + 0.5),
			Y:    int(v + 0.5),
		})
		y += t.EmHeight * t.LineHeight
	}

	return infos
}
