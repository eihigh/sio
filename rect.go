package sio

// anchor represents relative position like below:
// 7 8 9
// 4 5 6
// 1 2 3

// Rect is a simple rect
type Rect struct {
	x, y, w, h float64
}

// NewRect returns a new rect.
// Position is hidden, use Pos method.
func NewRect(anchor int, x, y, w, h float64) Rect {
	var r Rect
	r.Set(anchor, x, y, w, h)
	return r
}

// Set sets the data
func (r *Rect) Set(anchor int, x, y, w, h float64) {
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

// Move sets position relatively
func (r *Rect) Move(anchor int, x, y float64) {
	r.Set(anchor, x, y, r.w, r.h)
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

// CloneResizing returns a new resized rect, no breaking changes
func (r *Rect) CloneResizing(origAnchor, newAnchor int, diffX, diffY float64) *Rect {
	var ret Rect
	ret.w, ret.h = r.w+diffX, r.h+diffY
	x, y := r.Pos(origAnchor)
	ret.Move(newAnchor, x, y)
	return &ret
}

// CloneScaling returns a new scaled rect
func (r *Rect) CloneScaling(origAnchor, newAnchor int, scaleX, scaleY float64) *Rect {
	var ret Rect
	ret.w, ret.h = r.w*scaleX, r.h*scaleY
	x, y := r.Pos(origAnchor)
	ret.Move(newAnchor, x, y)
	return &ret
}
