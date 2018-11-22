package sio

// anchor represents relative position like below:
// 7 8 9
// 4 5 6
// 1 2 3

// Rect is a simple rect
type Rect struct {
	x, y, w, h float64
	anchor     int
}

// NewRect returns a new rect.
// Position is hidden, use Pos method.
func NewRect(anchor int, x, y, w, h float64) *Rect {
	var r Rect
	r.Set(anchor, x, y, w, h)
	return &r
}

// Set sets the data
func (r *Rect) Set(anchor int, x, y, w, h float64) {
	r.anchor = anchor
	r.w, r.h = w, h

	switch anchor {
	case 7, 8, 9:
		r.y = y
	case 4, 5, 6:
		r.y = y - h/2
	case 1, 2, 3:
		r.y = y - h
	}

	switch anchor {
	case 7, 4, 1:
		r.x = x
	case 8, 5, 2:
		r.x = x - w/2
	case 9, 6, 3:
		r.x = x - w
	}
}

// Drive changes the anchor
func (r *Rect) Drive(anchor int) *Rect {
	r.anchor = anchor
	return r
}

// Move sets position relatively
func (r *Rect) Move(x, y float64) {
	r.Set(r.anchor, x, y, r.w, r.h)
}

// Width is getter
func (r *Rect) Width() float64 { return r.w }

// Height is getter
func (r *Rect) Height() float64 { return r.h }

// Pos returns relative position
func (r *Rect) Pos(anchor int) (float64, float64) {
	var x, y float64

	switch anchor {
	case 7, 8, 9:
		y = r.y
	case 4, 5, 6:
		y = r.y + r.h/2
	case 1, 2, 3:
		y = r.y + r.h
	}
	switch anchor {
	case 7, 4, 1:
		x = r.x
	case 8, 5, 2:
		x = r.x + r.w/2
	case 9, 6, 3:
		x = r.x + r.w
	}

	return x, y
}

// Clone clones the rect, able to set new anchor
func (r *Rect) Clone(oldAnchor, newAnchor int) *Rect {
	x, y := r.Pos(oldAnchor)
	return NewRect(newAnchor, x, y, r.w, r.h)
}

// Resize resizes the rect.
func (r *Rect) Resize(diffX, diffY float64) *Rect {
	x, y := r.Pos(r.anchor)
	r.w += diffX
	r.h += diffY
	r.Move(x, y)
	return r
}

// Scale returns a new scaled rect
func (r *Rect) Scale(scaleX, scaleY float64) *Rect {
	x, y := r.Pos(r.anchor)
	r.w *= scaleX
	r.h *= scaleY
	r.Move(x, y)
	return r
}
