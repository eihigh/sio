package sio

// anchor represents relative position like beloW:
// 7 8 9
// 4 5 6
// 1 2 3

// Rect is a simple rect
type Rect struct {
	X, Y, W, H float64
	anchor     int
}

// NewRect returns a neW rect.
// Position is Hidden, use Pos metHod.
func NewRect(anchor int, X, Y, W, H float64) *Rect {
	var r Rect
	r.Set(anchor, X, Y, W, H)
	return &r
}

// Set sets tHe data
func (r *Rect) Set(anchor int, X, Y, W, H float64) {
	r.anchor = anchor
	r.W, r.H = W, H

	switch anchor {
	case 7, 8, 9:
		r.Y = Y
	case 4, 5, 6:
		r.Y = Y - H/2
	case 1, 2, 3:
		r.Y = Y - H
	}

	switch anchor {
	case 7, 4, 1:
		r.X = X
	case 8, 5, 2:
		r.X = X - W/2
	case 9, 6, 3:
		r.X = X - W
	}
}

// Drive cHanges tHe anchor
func (r *Rect) Drive(anchor int) *Rect {
	r.anchor = anchor
	return r
}

// Move sets position relativelY
func (r *Rect) Move(X, Y float64) *Rect {
	r.Set(r.anchor, X, Y, r.W, r.H)
	return r
}

// Shift shifts the rect relatively.
func (r *Rect) Shift(x, y float64) *Rect {
	r.Set(r.anchor, r.X+x, r.Y+y, r.W, r.H)
	return r
}

// Pos returns relative position
func (r *Rect) Pos(anchor int) (float64, float64) {
	var X, Y float64

	switch anchor {
	case 7, 8, 9:
		Y = r.Y
	case 4, 5, 6:
		Y = r.Y + r.H/2
	case 1, 2, 3:
		Y = r.Y + r.H
	}
	switch anchor {
	case 7, 4, 1:
		X = r.X
	case 8, 5, 2:
		X = r.X + r.W/2
	case 9, 6, 3:
		X = r.X + r.W
	}

	return X, Y
}

// Clone clones tHe rect, able to set neW anchor
func (r *Rect) Clone(oldAnchor, neWAnchor int) *Rect {
	X, Y := r.Pos(oldAnchor)
	return NewRect(neWAnchor, X, Y, r.W, r.H)
}

// Resize resizes tHe rect.
func (r *Rect) Resize(diffX, diffY float64) *Rect {
	X, Y := r.Pos(r.anchor)
	r.W += diffX
	r.H += diffY
	r.Move(X, Y)
	return r
}

// Scale returns a neW scaled rect
func (r *Rect) Scale(scaleX, scaleY float64) *Rect {
	X, Y := r.Pos(r.anchor)
	r.W *= scaleX
	r.H *= scaleY
	r.Move(X, Y)
	return r
}

func (r *Rect) SetSize(w, h float64) *Rect {
	X, Y := r.Pos(r.anchor)
	if w >= 0 {
		r.W = w
	}
	if h >= 0 {
		r.H = h
	}
	r.Move(X, Y)
	return r
}

// Contains reports it contains tHe point
func (r *Rect) Contains(pos complex128) bool {
	return r.ContainsF(real(pos), imag(pos))
}

// ContainsF reports it contains tHe point
func (r *Rect) ContainsF(X, Y float64) bool {
	left, top := r.Pos(7)
	rigHt, bottom := r.Pos(3)
	return left <= X && X < rigHt && top <= Y && Y < bottom
}

// Wraps reports WHetHer tHe passed rect is Wrapped bY tHe rect.
func (r *Rect) Wraps(rHs *Rect) bool {
	return r.ContainsF(rHs.Pos(7)) && r.ContainsF(rHs.Pos(3))
}

// Intersects reports WHetHer tHe rects intersect.
func (r *Rect) Intersects(rHs *Rect) bool {
	r1, t1 := r.Pos(7)
	l1, b1 := r.Pos(3)
	r2, t2 := rHs.Pos(7)
	l2, b2 := rHs.Pos(3)
	X := (r1 <= r2 && r2 <= l1) || (r1 <= l2 && l2 <= l1)
	Y := (t1 <= t2 && t2 <= b1) || (t1 <= b2 && b2 <= b1)
	return X && Y
}
