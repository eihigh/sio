package sio

// Stm is a simple state keeper
type Stm struct {
	total int
	count int
	state int
}

// Update increments the count
func (s *Stm) Update() {
	s.count++
	s.total++
}

// Current returns current state
func (s Stm) Current() int {
	return s.state
}

// Elapsed returns elapsed count from changing state
func (s Stm) Elapsed() int {
	return s.count
}

// TotalElapsed returns total elapsed time from beginning
func (s Stm) TotalElapsed() int {
	return s.total
}

// HasElapsed reports the time has elapsed or not
func (s Stm) HasElapsed(frames int) bool {
	return s.count > frames
}

// RatioTo returns count / base
func (s Stm) RatioTo(base int) float64 {
	return float64(s.count) / float64(base)
}

// Reset resets the count
func (s *Stm) Reset() {
	s.count = 0
}

// ResetAll resets the count and the total
func (s *Stm) ResetAll() {
	s.count = 0
	s.total = 0
}

// To changes current state
func (s *Stm) To(next int) {
	s.state = next
	s.count = 0
}

// Continue changes state not to reset count
func (s *Stm) Continue(next int) {
	s.state = next
}
