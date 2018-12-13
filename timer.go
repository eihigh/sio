package sio

type Timer struct {
	State        string
	Count, Limit int
}

func (t *Timer) Update() {
	if t.Limit != 0 && t.Count < t.Limit {
		t.Count++
		return
	}
	t.Count++
}

func (t *Timer) Switch(state string) {
	t.State = state
	t.Count = 0
}

func (t *Timer) Ratio() float64 {
	return float64(t.Count) / float64(t.Limit)
}

func (t *Timer) RatioTo(count int) float64 {
	return float64(t.Count) / float64(count)
}

func (t *Timer) Do(b, e int, f func(Timer)) (then Timer) {
	then = Timer{
		State: t.State,
		Count: t.Count - e,
		Limit: 0,
	}

	if b < 0 {
		b = t.Limit + b
	}
	if e <= 0 {
		e = t.Limit + e
	}

	if t.Count < b {
		return
	}
	if e <= t.Count {
		return
	}

	child := Timer{
		State: t.State,
		Count: t.Count - b,
		Limit: e - b,
	}
	f(child)
	return
}

func (t *Timer) Once(f func()) {
	if t.Count == 0 {
		f()
	}
}

type TimersMap map[string]*Timer

func (tm TimersMap) UpdateAll() {
	for _, t := range tm {
		t.Update()
	}
}
