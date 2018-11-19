package sio

var (
	// EmHeight is the height of 'M'
	EmHeight = 12
	// EmWidth is the width of 'M'
	EmWidth = 6

	// LetterSpacing is vertical spacing
	LetterSpacing = 1.0
	// LineHeight is horizontal spacing
	LineHeight = 1.5
)

// TextAlign is text drawing alignment
type TextAlign int

// TextAlign constants
const (
	AlignLeft TextAlign = iota
	AlignCenter
	AlignCenterAlways
	AlignRight
)

// TextBox provides some functions for monospace font drawing
type TextBox struct {
	align  TextAlign
	text   []rune
	cursor int
	rect   *Rect
	pos    complex128
}

// LineDrawInfo contains info to draw a line
type LineDrawInfo struct {
	Text string
	X, Y int
}

// SeekRune adds the last text position.
func (t *TextBox) SeekRune() {
	t.cursor++
}

// SeekLine adds the line.
func (t *TextBox) SeekLine() {
	for {
		if t.text[t.cursor] == '\n' {
			return
		}
		t.cursor++
	}
}

/*
info := textBox.DrawInfo()
for _, i := range info {
	text.Draw(scr, i.Text, fface, i.X, i.Y, color.White)
}

if tx.Elapsed() > 3 {
	textBox.SeekRune()
	tx.Reset()
}
textBox.Draw()
*/

func (t *TextBox) Draw() {

}
