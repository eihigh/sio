package sio

// Worker keeps count and state.
type Worker struct {
	Count int
	State string
}

// Switch switches the state and resets the count.
func (w *Worker) Switch(state string) {
	w.State = state
	w.Count = 0
}

// Continue changes the state without clearing the count.
func (w *Worker) Continue(state string) {
	w.State = state
}

// Do executes the function if the count is in the range of [b, e).
func (w *Worker) Do(b, e int, f func(t float64)) {
	if w.Count < b {
		return
	}
	if e <= w.Count {
		return
	}
	t := float64(w.Count-b) / float64(e-b)
	f(t)
}
