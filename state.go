package sio

// State is a simple state and a counter.
type State struct {
	state, count, age int
}

// Update increments the count.
func (s *State) Update() {
	s.count++
	s.age++
}

// Get returns the state.
func (s *State) Get() int {
	return s.state
}

// Count returns the count.
func (s *State) Count() int {
	return s.count
}

// Age returns the count from the beginning.
func (s *State) Age() int {
	return s.age
}

// HasCounted reports if the time has elapsed.
func (s *State) HasCounted(count int) bool {
	return s.count >= count
}

// HasAged reports if the time has elapsed from the beginning.
func (s *State) HasAged(count int) bool {
	return s.age >= count
}

// RatioTo returns the ratio of the count to the target.
func (s *State) RatioTo(target int) float64 {
	return float64(s.count) / float64(target)
}

// Reset resets the count.
func (s *State) Reset() {
	s.count = 0
}

// Rebirth resets the count and the age.
func (s *State) Rebirth() {
	s.count = 0
	s.age = 0
}

// To changes the state.
func (s *State) To(next int) {
	s.state = next
	s.count = 0
}

// Continue changes the state not to reset the count.
func (s *State) Continue(next int) {
	s.state = next
}
